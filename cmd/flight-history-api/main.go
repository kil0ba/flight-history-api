package main

import (
	"flag"

	"github.com/BurntSushi/toml"
	flighthistory "github.com/kil0ba/flight-history-api/internal/app/flight-history"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
)

var (
	configPath string
)

func init() {
	// Initing the env flags
	flag.StringVar(&configPath, "config-path", "configs/dev.toml", "The path for the config")
}

func main() {
	flag.Parse()
	config := flighthistoryserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		panic(err)
	}

	// TODO add logger
	server := &flighthistoryserver.FlightHistoryServer{
		Config: config,
	}

	flighthistory.Start(server)
}
