package config

type (
	Config struct {
		ServerConfig
	}

	ServerConfig struct {
		Port string
	}
)

func GetConfig() Config {
	return Config{
		ServerConfig: getServerConfig(),
	}
}

func getServerConfig() ServerConfig {
	return ServerConfig{
		Port: ":8080",
	}
}
