package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	// defaultAddressTimeout defines the duration to wait
	// for new addresses.
	defaultAddressTimeout = time.Minute * 10

	// defaultNodeTimeout defines the timeout time waiting for
	// a response from a node.
	defaultNodeTimeout = time.Second * 3
)

var (
	manager  *Manager
	globalWg sync.WaitGroup
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		os.Exit(1)
	}
	manager, err = NewManager(filepath.Join(defaultHomeDir,
		activeNetParams.Name))
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	manager.AddAddresses([]net.IP{net.ParseIP(cfg.Seeder)})

	globalWg.Add(1)
	go creep()

	dnsServer := NewDNSServer(cfg.Host, cfg.Nameserver, cfg.Listen)
	go dnsServer.Start()

	globalWg.Wait()
}
