package util

import "os"

// EnvConfig represents app configuration
type EnvConfig struct {
	Hostname   string
	Port       string
	GcpProject string
}

// LookupEnvWithDefault helper function that retrieves env variable
// and allows a fallback to default value
func LookupEnvWithDefault(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return fallback
}

// NewConfig creates new app config
func NewConfig() EnvConfig {
	return EnvConfig{
		Hostname:   LookupEnvWithDefault("HOSTNAME", "localhost"),
		Port:       LookupEnvWithDefault("PORT", "1317"),
		GcpProject: LookupEnvWithDefault("PROJECT_ID", "test"),
	}
}
