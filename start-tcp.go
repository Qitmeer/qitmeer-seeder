package main

import (
	"fmt"
	"github.com/HalalChain/qitmeer/core/protocol"
	"github.com/miekg/dns"
	"log"
	"net"
	"strconv"
	"strings"
)

func (d *DNSServer) StartTCP() {
	defer wg.Done()

	rr := fmt.Sprintf("%s 86400 IN NS %s", d.hostname, d.nameserver)
	authority, err := dns.NewRR(rr)
	if err != nil {
		log.Printf("NewRR: %v", err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", d.listen)
	if err != nil {
		log.Printf("ResolveTCPAddr: %v", err)
		return
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Printf("ListenTCP: %v", err)
		return
	}
	//noinspection GoUnhandledErrorResult
	defer tcpListener.Close()

	for {
		b := make([]byte, 512)
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Printf("Read: %v", err)
			continue
		}

		go func() {
			dnsMsg := new(dns.Msg)
			err = dnsMsg.Unpack(b[:])
			if err != nil {
				log.Printf("%s: invalid dns message: %v",
					tcpConn, err)
				return
			}
			if len(dnsMsg.Question) != 1 {
				log.Printf("%s sent more than 1 question: %d",
					tcpConn, len(dnsMsg.Question))
				return
			}
			domainName := strings.ToLower(dnsMsg.Question[0].Name)
			ff := strings.LastIndex(domainName, d.hostname)
			if ff < 0 {
				log.Printf("invalid name: %s",
					dnsMsg.Question[0].Name)
				return
			}

			wantedSF := protocol.Full
			labels := dns.SplitDomainName(domainName)
			if labels[0][0] == 'x' && len(labels[0]) > 1 {
				wantedSFStr := labels[0][1:]
				u, err := strconv.ParseUint(wantedSFStr, 10, 64)
				if err != nil {
					log.Printf("%s: ParseUint: %v", tcpConn, err)
					return
				}
				wantedSF = protocol.ServiceFlag(u)
			}

			var atype string
			qtype := dnsMsg.Question[0].Qtype
			switch qtype {
			case dns.TypeA:
				atype = "A"
			case dns.TypeAAAA:
				atype = "AAAA"
			case dns.TypeNS:
				atype = "NS"
			default:
				log.Printf("%s: invalid qtype: %d", tcpConn,
					dnsMsg.Question[0].Qtype)
				return
			}

			log.Printf("%s: query %d for %v", tcpConn,
				dnsMsg.Question[0].Qtype, wantedSF)

			respMsg := dnsMsg.Copy()
			respMsg.Authoritative = true
			respMsg.Response = true

			if qtype != dns.TypeNS {
				respMsg.Ns = append(respMsg.Ns, authority)
				ips := manager.GoodAddresses(qtype, wantedSF)
				for _, ip := range ips {
					rr = fmt.Sprintf("%s 30 IN %s %s",
						dnsMsg.Question[0].Name, atype,
						ip.String())
					newRR, err := dns.NewRR(rr)
					if err != nil {
						log.Printf("%s: NewRR: %v",
							tcpConn, err)
						return
					}

					respMsg.Answer = append(respMsg.Answer,
						newRR)
				}
			} else {
				rr = fmt.Sprintf("%s 86400 IN NS %s",
					dnsMsg.Question[0].Name, d.nameserver)
				newRR, err := dns.NewRR(rr)
				if err != nil {
					log.Printf("%s: NewRR: %v", tcpConn, err)
					return
				}

				respMsg.Answer = append(respMsg.Answer, newRR)
			}

			//done:
			sendBytes, err := respMsg.Pack()
			if err != nil {
				log.Printf("%s: failed to pack response: %v",
					tcpConn, err)
				return
			}

			_, err = tcpConn.Write(sendBytes)
			if err != nil {
				log.Printf("%s: failed to write response: %v",
					tcpConn, err)
				return
			}
		}()
	}
}
