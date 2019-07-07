package main

import (
	"github.com/HalalChain/qitmeer-lib/core/dag"
	"github.com/HalalChain/qitmeer-lib/core/message"
	"github.com/HalalChain/qitmeer/p2p/peer"
	"log"
	"net"
	"sync"
	"time"
)

func creep() {
	defer globalWg.Done()

	onaddr := make(chan struct{})
	verack := make(chan struct{})

	NewestGSFunc := func() (gs *dag.GraphState, err error) {
		gs = dag.NewGraphState()
		gs.GetTips().Add(activeNetParams.GenesisHash)
		gs.SetTotal(1)
		return
	}

	config := peer.Config{
		UserAgentName:    "hlc-seeder",
		UserAgentVersion: "0.0.2",
		ChainParams:      activeNetParams,
		DisableRelayTx:   true,

		NewestGS: NewestGSFunc,

		Listeners: peer.MessageListeners{
			OnAddr: func(p *peer.Peer, msg *message.MsgAddr) {
				n := make([]net.IP, 0, len(msg.AddrList))
				for _, addr := range msg.AddrList {
					n = append(n, addr.IP)
				}
				added := manager.AddAddresses(n)
				log.Printf("Peer %v sent %v addresses, %d new",
					p.Addr(), len(msg.AddrList), added)
				onaddr <- struct{}{}
			},
			OnVerAck: func(p *peer.Peer, msg *message.MsgVerAck) {
				log.Printf("Adding peer %v with services %v",
					p.NA().IP.String(), p.Services())

				verack <- struct{}{}
			},
		},
	}

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

				host := net.JoinHostPort(ip.String(),
					activeNetParams.DefaultPort)
				p, err := peer.NewOutboundPeer(&config, host)
				if err != nil {
					log.Printf("NewOutboundPeer on %v: %v",
						host, err)
					return
				}
				manager.Attempt(ip)
				conn, err := net.DialTimeout("tcp", p.Addr(),
					defaultNodeTimeout)
				if err != nil {
					return
				}
				p.AssociateConnection(conn)

				// Wait for the verack message or timeout in case of
				// failure.
				select {
				case <-verack:
					// Mark this peer as a good node.
					manager.Good(p.NA().IP, p.Services())

					// Ask peer for some addresses.
					p.QueueMessage(message.NewMsgGetAddr(), nil)

				case <-time.After(defaultNodeTimeout):
					log.Printf("verack timeout on peer %v",
						p.Addr())
					p.Disconnect()
					return
				}

				select {
				case <-onaddr:
				case <-time.After(defaultNodeTimeout):
					log.Printf("getaddr timeout on peer %v",
						p.Addr())
					p.Disconnect()
					return
				}
				p.Disconnect()
			}(ip)
		}
		wg.Wait()
	}
}
