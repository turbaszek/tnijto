package utils

import "os"

// EnvConfig represents app configuration
type EnvConfig struct {
	Hostname string
}

// NewConfig creates new app config
func NewConfig() EnvConfig {
	hostname, exists := os.LookupEnv("HOSTNAME")
	if !exists {
		hostname = "localhost"
	}
	return EnvConfig{
		Hostname: hostname,
	}
}
