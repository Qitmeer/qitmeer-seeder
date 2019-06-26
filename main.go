package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
)

var (
	manager *Manager
	wg      sync.WaitGroup
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Println("loadConfig : ", err.Error())
		os.Exit(1)
	}

	manager, err = NewManager(filepath.Join(defaultHomeDir, activeNetParams.Name))
	if err != nil {
		log.Println("NewManager : ", err.Error())
		os.Exit(1)
	}

	manager.AddAddresses([]net.IP{net.ParseIP(cfg.Seeder)})

	wg.Add(1)

	go creep()

	dnsServer := NewDNSServer(cfg.Host, cfg.NameServer, cfg.Listen)

	go dnsServer.StartTCP()

	wg.Wait()
}
