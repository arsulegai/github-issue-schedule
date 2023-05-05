package utils

import "os"

// GetEnvOrDefault returns the default value for an environment
// variable if the value is not explicitly set
func GetEnvOrDefault(env, defaultValue string) string {
	value, isPresent := os.LookupEnv(env)
	if isPresent {
		return value
	}
	return defaultValue
}
