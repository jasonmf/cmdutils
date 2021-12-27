// Package env has utility functions for retrieving values from environment
// variables.
package env

import (
	"log"
	"os"
)

// Must retrieves a value from an env var and returns it as a string. If the
// resulting string is empty Must will call log.Fatalf with a helpful message
func Must(envVar string) string {
	v := os.Getenv(envVar)
	if v == "" {
		log.Fatalf("required environment variable %s empty", envVar)
	}
	return v
}
