package serverConfig

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	DebugLevel string `toml:"debug_level"`
	Db         string `toml:"db"`
	Secret     string `toml:"secret"`
	JwtSecret  string `toml:"jwtSecret"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		DebugLevel: "",
		Secret:     "",
		JwtSecret:  "",
	}
}
