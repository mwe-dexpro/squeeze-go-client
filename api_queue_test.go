package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueApi_GetQueueSteps(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	steps, e := c.Queue.GetQueueSteps()
	assert.Nil(t, e)
	assert.NotNil(t, steps)
	assert.NotEmpty(t, steps)
}

func TestQueueApi_GetQueueStep(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	step, e := c.Queue.GetQueueStep("Validation")
	assert.Nil(t, e)
	assert.NotNil(t, step)
	assert.NotEmpty(t, step)
	assert.Equal(t, "Validation", step.Name)
	assert.NotEmpty(t, step.StepDetails)
}
