package utility

import (
	"os"
)

// Get values of environment variables(env).
// These environment variables are defined in docker-compose.yml file
// If the env exists, function returns it's value, otherwise, returns default value
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}
	return defaultValue
}