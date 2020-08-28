package util

import "os"

// EnvConfig represents app configuration
type EnvConfig struct {
	Hostname   string
	Port       string
	GcpProject string
}

func lookupEnvWithDefault(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

func newConfig() EnvConfig {
	return EnvConfig{
		Hostname:   lookupEnvWithDefault("HOSTNAME", "localhost"),
		Port:       lookupEnvWithDefault("PORT", "1317"),
		GcpProject: lookupEnvWithDefault("PROJECT_ID", "test"),
	}
}

// Config represents current app configuration
var Config = newConfig()
