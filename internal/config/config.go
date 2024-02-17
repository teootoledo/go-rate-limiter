package config

import "os"

type Configuration struct {
	ServerPort string
}

const (
	ServerPort        = "SERVER_PORT"
	DefaultServerPort = "8080"
)

var config *Configuration

func GetConfig() *Configuration {
	if config == nil {
		config = initConfig()
	}

	return config
}

func initConfig() *Configuration {
	return &Configuration{
		ServerPort: getEnvOrDefault(ServerPort, DefaultServerPort),
	}
}

func getEnvOrDefault(environmentVarName string, defValue string) string {
	if environmentVar := os.Getenv(environmentVarName); environmentVar != "" {
		return environmentVar
	}

	return defValue
}
