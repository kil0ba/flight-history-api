package flighthistoryserver

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	DebugLevel string `toml:"debug_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		DebugLevel: "",
	}
}
