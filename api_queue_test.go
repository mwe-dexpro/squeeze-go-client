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

func TestQueueApi_GetQueueStep(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = getEnvApiKey(t)
	step, e := c.Queue.GetQueueStep("Validation")
	assert.Nil(t, e)
	assert.NotNil(t, step)
	assert.NotEmpty(t, step)
	assert.Equal(t, "Validation", step.Name)
}
