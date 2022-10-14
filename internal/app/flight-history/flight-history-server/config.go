package flighthistoryserver

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	DebugLevel string `toml:"debug_level"`
	Db         string `toml:"db"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		DebugLevel: "",
	}
}
