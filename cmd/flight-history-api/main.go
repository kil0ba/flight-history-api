package main

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	flighthistory "github.com/kil0ba/flight-history-api/internal/app/flight-history"
	flighthistoryserver "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server"
	serverConfig "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
)

var (
	configPath string
)

func init() {
	isDebug := os.Getenv("debug")
	var path string

	// For VS Code
	if isDebug == "true" {
		path = "../../configs/dev.toml"
	} else {
		path = "configs/dev.toml"
	}
	// Initing the env flags
	flag.StringVar(&configPath, "config-path", path, "The path for the config")
}

func main() {
	flag.Parse()
	config := serverConfig.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		panic(err)
	}

	server := flighthistoryserver.New(config)

	flighthistory.Start(server)
}
