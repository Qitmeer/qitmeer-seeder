package main

import (
	"fmt"
	"github.com/HalalChain/qitmeer-lib/core/protocol"
	"github.com/HalalChain/qitmeer-lib/params"
	"github.com/jessevdk/go-flags"
	"log"
	"net"
	"os"
)

const (
	defaultListenPort = "18130"
	defaultSeed       = "seed.fulingjie.com"
)

var (
	// Default network parameters
	activeNetParams = &params.TestNetParams

	// Default configuration options
	defaultHomeDir, _ = os.Getwd()
)

// config defines the configuration options for hardforkdemo.
//
// See loadConfig for details on the configuration load process.
type config struct {
	Host       string `short:"H" long:"host" description:"Seed DNS address"`
	Listen     string `short:"l" long:"listen"  description:"Listen on address:port"`
	Nameserver string `short:"n" long:"nameserver" description:"hostname of nameserver"`
	Seeder     string `short:"s" long:"default seeder" description:"IP address of a  working node"`
	TestNet    bool   `short:"t" long:"testnet" description:"Use testnet"`
}

func loadConfig() (*config, error) {

	log.SetPrefix("hlc-seeder")
	log.SetFlags(log.Lshortfile | log.Ldate)

	err := os.MkdirAll(defaultHomeDir, os.ModePerm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Default config.
	cfg := config{
		Listen: normalizeAddress("0.0.0.0", defaultListenPort),
	}

	preCfg := cfg
	preParser := flags.NewParser(&preCfg, flags.Default)
	_, err = preParser.Parse()
	if err != nil {
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			os.Exit(0)
		}
		preParser.WriteHelp(os.Stderr)
		return nil, err
	}
	// Load additional config from file.
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
		err := fmt.Errorf(str)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	if len(cfg.Nameserver) == 0 {
		str := "Please specify a nameserver"
		err := fmt.Errorf(str)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	if len(cfg.Seeder) == 0 {
		str := "Please specify a seeder"
		err := fmt.Errorf(str)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return nil, err
	}

	cfg.Listen = normalizeAddress(cfg.Listen, defaultListenPort)

	if cfg.TestNet {
		activeNetParams = &params.TestNetParams
		activeNetParams.DNSSeeds = []params.DNSSeed{
			{defaultSeed, true},
			{defaultSeed, true},
			{defaultSeed, true},
		}
		activeNetParams.Net = protocol.TestNet
		activeNetParams.DefaultPort = defaultListenPort
	}

	return &cfg, nil
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
