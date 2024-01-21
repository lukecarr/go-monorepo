package osx

import (
	"os"
	"strings"
)

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will default to the provided default (`def`)
// if the variable isn't found.
func Getenv(key, def string) string {
	if val, ok := os.LookupEnv(key); ok && len(strings.TrimSpace(val)) > 0 {
		return val
	}

	return def
}
