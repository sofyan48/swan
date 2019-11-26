package libs

import (
	"os"
)

// GetEnvVariabel Getting Variabel From Environment
func GetEnvVariabel(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
