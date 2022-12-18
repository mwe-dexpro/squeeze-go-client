package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDocumentApi_ProcessDocument(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)

	filepath := "testdata/brt.pdf"
	filename := "brt.pdf"

	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			// As long as we do not version testdata in this git repository, we skip this test in the CI
			t.Skip("Test file not found: " + filepath)
		} else {
			t.Error(err)
		}
	}

	file, err := os.Open(filepath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	classes, e := c.Document.ProcessDocument(1, 0, "", nil, file, filename)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}

func TestDocumentApi_GetDocumentById(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)

	classes, e := c.Document.GetDocumentById(1)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}
