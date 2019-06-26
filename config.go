package main

import (
	"github.com/HalalChain/qitmeer/params"
	"log"
	"os"
)

var (
	// Default network parameters
	activeNetParams = &params.TestNetParams
	//get current path
	defaultHomeDir, _ = os.Getwd()
)

// See loadConfig for details on the configuration load process.
type config struct {
	Host       string
	Listen     string
	NameServer string
	Seeder     string
}

// TODO: Add TOML configuration support
func loadConfig() (*config, error) {
	err := os.MkdirAll(defaultHomeDir, os.ModePerm)
	if err != nil {
		log.Fatalln("MkdirAll : ", err.Error())
	}

	//Init config
	cfg := config{
		Host:       "seeder.fulingjie.com",
		Listen:     "0.0.0.0:18130",
		NameServer: "seeder.fulingjie.com",
		Seeder:     "118.24.234.251",
	}

	//Change seeds
	activeNetParams.DNSSeeds =[]params.DNSSeed{
		{"seeder.fulingjie.com", true},
		{"seeder.fulingjie.com", true},
		{"seeder.fulingjie.com", true},
	}

	return &cfg, nil
}
