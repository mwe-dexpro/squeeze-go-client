package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublicApi_GetHealth(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	e := c.Public.GetHealth()
	assert.Nil(t, e)
}
