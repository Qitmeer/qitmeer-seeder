package main

import (
	"log"
	"net"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/Qitmeer/qitmeer/params"
)

var defaultListenPort = params.MainNetParams.DefaultPort //"18130"

// Default network parameters
var activeNetParams = &params.MainNetParams

// Default configuration options
var defaultHomeDir, _ = os.Getwd()

type config struct {
	Host       string `short:"H" long:"host" description:"Seed DNS address"`
	Listen     string `short:"l" long:"listen"  description:"Listen on address:port"`
	Nameserver string `short:"n" long:"nameserver" description:"hostname of nameserver"`
	Seeder     string `short:"s" long:"default seeder" description:"IP address of a  working node"`
	TestNet    bool   `short:"t" long:"testnet" description:"Use testnet"`
}

func loadConfig() (*config, error) {
	log.SetPrefix("qitmeer-seeder")
	log.SetFlags(log.Lshortfile | log.Ldate)

	err := os.MkdirAll(defaultHomeDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Default config.
	cfg := config{
		Listen: normalizeAddress("0.0.0.0", defaultListenPort),
	}

	parser := flags.NewParser(&cfg, flags.Default)

	// Parse command line options again to ensure they take precedence.
	_, err = parser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); !ok || e.Type != flags.ErrHelp {
			parser.WriteHelp(os.Stderr)
		}
		return nil, err
	}

	if len(cfg.Host) == 0 {
		str := "Please specify a hostname"
		log.Fatalln(str)
	}

	if len(cfg.Nameserver) == 0 {
		str := "Please specify a nameserver"
		log.Fatalln(str)
	}

	if len(cfg.Seeder) == 0 {
		str := "Please specify a seeder"
		log.Fatalln(str)
	}

	cfg.Listen = normalizeAddress(cfg.Listen, defaultListenPort)

	if cfg.TestNet {
		activeNetParams = &params.TestNetParams
		defaultListenPort = params.TestNetParams.DefaultPort
	}

	return &cfg, err
}

// normalizeAddress returns addr with the passed default port appended if
// there is not already a port specified.
func normalizeAddress(addr, defaultPort string) string {
	_, _, err := net.SplitHostPort(addr)
	if err != nil {
		return net.JoinHostPort(addr, defaultPort)
	}
	return addr
}
