package squeeze_go_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicApi_GetHealth(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	e := c.Public.GetHealth()
	assert.Nil(t, e)
}
