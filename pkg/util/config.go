package util

import "os"

// EnvConfig represents app configuration
type EnvConfig struct {
	// Hostname is used to construct redirect link
	// For example http://{Hostname}/happiness
	Hostname string

	// Port is used to construct redirect link
	// For example http://{Hostname}:{Port}/happiness
	Port string

	// GcpProject should have the value of Google Cloud project
	// that will be used to store data in Firestore
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
