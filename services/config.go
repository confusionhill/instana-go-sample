package services

type Config struct {
	JwtSecret string
}

var cfg *Config

func GetConfig() *Config {
	if cfg == nil {
		LoadConfig()
	}
	return cfg
}

func LoadConfig() {
	cfg = &Config{
		JwtSecret: "secret-1-2-3-4",
	}
}
