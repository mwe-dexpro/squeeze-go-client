package squeeze_go_client

import (
	"fmt"
	"os"
	"testing"
)

// getEnvVal either returns the value of the ENV var with the given name or panics if it is not set.
func getEnvVal(key string) string {
	if val, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("environment variable %s is not set", key))
	} else {
		return val
	}
}

// getEnvApiKey either returns the value of the ENV var SQZ_KEY or skips the test if it is not set.
func getEnvApiKey(t *testing.T) string {
	if val, ok := os.LookupEnv("SQZ_KEY"); !ok {
		t.Skip("environment variable SQZ_KEY is not set")
		return ""
	} else {
		return val
	}
}
