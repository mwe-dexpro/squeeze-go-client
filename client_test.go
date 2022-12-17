package squeeze_go_client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getEnvVal(key string) string {
	if val, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("environment variable %s is not set", key))
	} else {
		return val
	}
}

func TestPublicApi_GetHealth(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	e := c.Public.GetHealth()
	assert.Nil(t, e)
}
