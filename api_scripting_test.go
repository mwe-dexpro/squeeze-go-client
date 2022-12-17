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

func TestScriptingApi_ExecuteScript(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	scripts, e := c.ScriptingApi.GetScripts()

	assert.Nil(t, e)
	assert.NotNil(t, scripts)
	assert.NotEmpty(t, scripts)

	// Get unlock script, that one is safe to test in CI / etc.
	var unlockScript *ScriptDto
	for _, script := range scripts {
		if script.Name == "(Internal) unlock-documents" {
			unlockScript = script
			break
		}
	}

	if unlockScript == nil {
		t.Error("Could not find unlock script to test execution with")
	}

	e = c.ScriptingApi.ExecuteScript(unlockScript.Id, false)
	assert.Nil(t, e)
}
