package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScriptingApi_GetScripts(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	scripts, e := c.ScriptingApi.GetScripts()
	assert.Nil(t, e)
	assert.NotNil(t, scripts)
	assert.NotEmpty(t, scripts)
}
