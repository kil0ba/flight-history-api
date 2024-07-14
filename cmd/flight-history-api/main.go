package main

import (
	"context"
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

//	@title			Flight History API
//	@version		0.1
//	@description	Flight History API documentation.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

// @securitydefinitions.oauth2.application	OAuth2Application
// @tokenUrl								https://example.com/oauth/token
// @scope.write							Grants write access
// @scope.admin							Grants read and write access to administrative information
func main() {
	flag.Parse()
	config := serverConfig.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	server := flighthistoryserver.New(ctx, config)

	flighthistory.Start(ctx, server)
}
