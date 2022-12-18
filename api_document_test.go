package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentApi_ProcessDocument(t *testing.T) {
	// todo:
}

func TestDocumentApi_GetDocumentById(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)

	classes, e := c.Document.GetDocumentById(1)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}
