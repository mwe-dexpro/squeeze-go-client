package squeeze_go_client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentClassApi_GetDocumentClasses(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = getEnvApiKey(t)
	classes, e := c.DocumentClass.GetDocumentClasses()
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}

func TestDocumentClassApi_GetAllFieldGroups(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = getEnvApiKey(t)
	classes, e := c.DocumentClass.GetAllFieldGroups(1)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}

func TestDocumentClassApi_GetAllDocumentClassFields(t *testing.T) {
	c := NewClient(getEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = getEnvApiKey(t)
	classes, e := c.DocumentClass.GetAllDocumentClassFields(1)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}
