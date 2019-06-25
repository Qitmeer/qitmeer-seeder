package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
	"net"
)

type DNSServer struct {
	hostname   string
	listen     string
	nameserver string
}

func (dnsServer *DNSServer) Start() {
	defer wg.Done()

	rrStr := fmt.Sprintf("%s 86400 IN NS %s", dnsServer.hostname, dnsServer.nameserver)
	authority, err := dns.NewRR(rrStr)
	if err != nil {
		log.Printf("NewRR: %v", err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", dnsServer.listen)
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
		conn, err := tcpListener.Accept()

		if err != nil {
			log.Printf("Accept: %v", err)
			continue
		}

		go dnsServer.handleConnection(rrStr, authority, conn)
	}
}

func NewDNSServer(hostname, nameserver, listen string) *DNSServer {
	if hostname[len(hostname)-1] != '.' {
		hostname = hostname + "."
	}
	if nameserver[len(nameserver)-1] != '.' {
		nameserver = nameserver + "."
	}

	return &DNSServer{
		hostname:   hostname,
		listen:     listen,
		nameserver: nameserver,
	}
}
