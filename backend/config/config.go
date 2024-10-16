// config.go - Loads environment variables for the application.
package config

import (
	"log"
	"os"
)

func LoadEnv() {
	// Example environment variable setup
	os.Setenv("SOURCEGRAPH_API_URL", "https://sourcegraph.com/api") // Replace with actual environment setup
	log.Println("Environment variables loaded")
}
