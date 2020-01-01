package main

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/Qitmeer/qitmeer/core/blockdag"
	"github.com/Qitmeer/qitmeer/core/message"
	"github.com/Qitmeer/qitmeer/core/protocol"
	"github.com/Qitmeer/qitmeer/p2p/connmgr"
	"github.com/Qitmeer/qitmeer/p2p/peer"
)

var checkGSMap map[uint]*blockdag.GraphState
var checkMaxMainHeight uint

//
var NewestGSFunc peer.GSFunc

func init() {
	GenesisCheckGS := blockdag.NewGraphState()
	GenesisCheckGS.GetTips().Add(activeNetParams.GenesisHash)
	GenesisCheckGS.SetTotal(1)

	NewestGS := func() (*blockdag.GraphState, error) {
		return GenesisCheckGS, nil
	}
	NewestGSFunc = NewestGS

	checkGSMap = make(map[uint]*blockdag.GraphState)
	checkGSMap[0] = GenesisCheckGS
}

func creep() {
	defer globalWg.Done()

	var wg sync.WaitGroup
	for {
		ips := manager.Addresses()
		if len(ips) == 0 {
			log.Printf("No stale addresses -- sleeping for %v",
				defaultAddressTimeout)
			time.Sleep(defaultAddressTimeout)
			continue
		}

		wg.Add(len(ips))

		for _, ip := range ips {
			go func(ip net.IP) {
				defer wg.Done()

				host := net.JoinHostPort(ip.String(), activeNetParams.DefaultPort)

				//p, err := peer.NewOutboundPeer(&peerConfig, host)
				sp := NewServerPeer(host)
				if sp == nil {
					return
				}
				p := sp.P

				manager.Attempt(ip)
				conn, err := net.DialTimeout("tcp", p.Addr(), defaultNodeTimeout)
				if err != nil {
					return
				}

				c := connmgr.NewConnReq()
				c.SetConn(conn)

				p.AssociateConnection(c)

				// Wait for the verack message or timeout in case of
				// failure.
				select {
				case isBad := <-sp.IsBadPeerCh:
					log.Printf("peer: %s isBad: %v ", p.NA().IP.String(), isBad)
					if isBad {
						manager.Bad(p.NA().IP)
						p.Disconnect()
						return
					}
					p.QueueMessage(message.NewMsgGetAddr(), nil)
				case <-time.After(defaultNodeTimeout):
					log.Printf("isBad timeout on peer %v", p.NA().IP.String())
					manager.Bad(p.NA().IP)
					p.Disconnect()
					return
				}

				select {
				case <-sp.OnAddrCh:
				case <-time.After(defaultNodeTimeout):
					log.Printf("getaddr timeout on peer %v", p.NA().IP.String())
					p.Disconnect()
					return
				}
				p.Disconnect()
			}(ip)
		}
		wg.Wait()
	}
}

// ServerPeer servefr peer
type ServerPeer struct {
	OnAddrCh chan struct{}
	// VerackCh    chan struct{}
	IsBadPeerCh chan bool

	P *peer.Peer

	checkGS *blockdag.GraphState
}

// NewServerPeer peer server
func NewServerPeer(host string) *ServerPeer {
	sp := &ServerPeer{
		OnAddrCh: make(chan struct{}),
		// VerackCh:    make(chan struct{}),
		IsBadPeerCh: make(chan bool),
	}

	messageListener := peer.MessageListeners{
		OnVersion: sp.OnVersion,
		OnAddr:    sp.OnAddrFunc,
		//OnVerAck:  sp.OnVerAckFunc,
		OnBlock: sp.OnBlock,
	}

	peerConfig := peer.Config{
		NewestGS:          NewestGSFunc,
		UserAgentName:     "qitmeer-seeder",
		UserAgentVersion:  "0.3.1",
		UserAgentComments: []string{"qitmeer", "seeder"},
		ChainParams:       activeNetParams,
		DisableRelayTx:    true,
		Services:          protocol.Full,
		ProtocolVersion:   protocol.ProtocolVersion,
		Listeners:         messageListener,
	}

	p, err := peer.NewOutboundPeer(&peerConfig, host)
	if err != nil {
		log.Printf("NewOutboundPeer on %v: %v",
			host, err)
		return nil
	}
	sp.P = p

	return sp

}

// OnAddrFunc on addr msg
func (sp *ServerPeer) OnAddrFunc(p *peer.Peer, msg *message.MsgAddr) {
	n := make([]net.IP, 0, len(msg.AddrList))
	for _, addr := range msg.AddrList {
		n = append(n, addr.IP)
	}
	added := manager.AddAddresses(n)
	log.Printf("Peer %v sent %v addresses, %d new",
		p.Addr(), len(msg.AddrList), added)

	// for _, addrrr := range msg.AddrList {
	// 	log.Println(p.Addr(), "add", addrrr.IP.String())
	// }

	sp.OnAddrCh <- struct{}{}
}

// func (sp *ServerPeer) OnVerAckFunc(p *peer.Peer, msg *message.MsgVerAck) {
// 	log.Printf("%s,Adding peer %v with services %v", p.Addr(),
// 		p.NA().IP.String(), p.Services())
// 	sp.VerackCh <- struct{}{}
// }

// OnBlock on block msg
func (sp *ServerPeer) OnBlock(p *peer.Peer, msg *message.MsgBlock, buf []byte) {
	//
	if sp.checkGS == nil {
		sp.IsBadPeerCh <- true
		return
	}

	if msg.Block.BlockHash().String() != sp.checkGS.GetMainChainTip().String() {
		log.Printf("onBlock,checkBlockHash not eq: peer: %v,getHash: %s,checkHash: %s", p.NA().IP,
			msg.Block.BlockHash().String(), sp.checkGS.GetMainChainTip().String())

		sp.IsBadPeerCh <- true
		return
	}
	sp.IsBadPeerCh <- false
}

// OnVersion on version msg
func (sp *ServerPeer) OnVersion(p *peer.Peer, msg *message.MsgVersion) *message.MsgReject {
	mainHeight := msg.LastGS.GetMainHeight()

	if p.NA().IP.String() == cfg.Seeder {
		log.Println("onVersion: seeder peer,good and update checkGS")

		//todo
		//mainHeight<1000
		if mainHeight > checkMaxMainHeight+10 {
			log.Println("add checkGS", mainHeight)
			checkGSMap[mainHeight] = msg.LastGS
			checkMaxMainHeight = mainHeight

			if len(checkGSMap) > 3 {
				var checkMainHeightMin uint = checkMaxMainHeight
				for checkMainHeight := range checkGSMap {
					if checkMainHeight < checkMainHeightMin {
						checkMainHeightMin = checkMainHeight
					}
				}
				delete(checkGSMap, checkMainHeightMin)
			}
		}
		sp.IsBadPeerCh <- false
		return nil
		//message.NewMsgReject(msg.Command(), message.RejectDuplicate, "duplicate peer version message")
	}

	log.Printf("onVersion:ip: %v,mainHeight: %d,order: %d,mainTip: %v", p.NA().IP, msg.LastGS.GetMainHeight(), msg.LastGS.GetMainOrder(), msg.LastGS.GetMainChainTip())

	if mainHeight == 0 {
		sp.IsBadPeerCh <- true
		return nil
	}

	if mainHeight == checkMaxMainHeight {
		if checkGSMap[checkMaxMainHeight].GetMainChainTip().String() == msg.LastGS.GetMainChainTip().String() {
			log.Println("dirt ok")
			sp.IsBadPeerCh <- false
			return nil
		}
	}

	for heightKey, gsValue := range checkGSMap {
		if mainHeight >= heightKey {
			sp.checkGS = gsValue
			break
		}
	}
	if sp.checkGS == nil {
		sp.IsBadPeerCh <- true
		return nil
	}

	//
	msgResp := message.NewMsgGetData()
	err := msgResp.AddInvVect(message.NewInvVect(message.InvTypeBlock, sp.checkGS.GetMainChainTip()))
	if err != nil {
		log.Println("onVersion: msgResp err: ", err, sp.checkGS.GetMainChainTip())
		sp.IsBadPeerCh <- true
		return message.NewMsgReject(msg.Command(), message.RejectInvalid, "")
	}
	p.QueueMessage(msgResp, nil)
	return nil
}
