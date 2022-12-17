package squeeze_go_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueApi_GetQueueSteps(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = getEnvApiKey(t)
	steps, e := c.Queue.GetQueueSteps()
	assert.Nil(t, e)
	assert.NotNil(t, steps)
	assert.NotEmpty(t, steps)
}
