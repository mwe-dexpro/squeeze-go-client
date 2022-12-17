package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJobsApi_GetAllJobs(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	scripts, e := c.Jobs.GetAllJobs()
	assert.Nil(t, e)
	assert.NotNil(t, scripts)
	assert.NotEmpty(t, scripts)
}
