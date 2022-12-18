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

func TestQueueApi_GetQueueStepEntriesSimple(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	response, e := c.Queue.GetQueueStepEntriesSimple("Validation", "", 0, 0, 1, 25)
	assert.Nil(t, e)
	assert.NotNil(t, response)

	assert.Equal(t, 25, response.Pagination.PageSize)
	assert.NotNil(t, response.Elements)
}

func TestQueueApi_RequeueDocument(t *testing.T) {
	t.Skip("Requeueing a document is for now disabled so that CI tests don't spam our test system")

	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	e := c.Queue.RequeueDocument(1, 0, "Extraction")
	assert.Nil(t, e)
}
