package squeeze_go_client

import (
	"github.com/dexpro-solutions-gmbh/squeeze-go-client/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocumentClassApi_GetDocumentClasses(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	classes, e := c.DocumentClass.GetDocumentClasses()
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}

func TestDocumentClassApi_GetAllFieldGroups(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	classes, e := c.DocumentClass.GetAllFieldGroups(1)
	assert.Nil(t, e)
	assert.NotNil(t, classes)
	assert.NotEmpty(t, classes)
}

func TestDocumentClassApi_GetAllDocumentClassFields(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)
	fields, e := c.DocumentClass.GetAllDocumentClassFields(1)
	assert.Nil(t, e)
	assert.NotNil(t, fields)
	assert.NotEmpty(t, fields)
}

// TestDocumentClassApi_Tables tests multiple endpoints related to tables
func TestDocumentClassApi_Tables(t *testing.T) {
	c := NewClient(internal.GetEnvVal("SQZ_BASE_PATH"))
	c.ApiKey = internal.GetEnvApiKey(t)

	tableId := 0

	t.Run("GetAllDocumentClassTables", func(t *testing.T) {
		tables, e := c.DocumentClass.GetAllDocumentClassTables(1)
		assert.Nil(t, e)
		assert.NotNil(t, tables)
		assert.NotEmpty(t, tables)

		tableId = tables[0].Id
	})

	t.Run("GetAllDocumentClassTableColumns", func(t *testing.T) {
		columns, e := c.DocumentClass.GetAllDocumentClassTableColumns(1, tableId)
		assert.Nil(t, e)
		assert.NotNil(t, columns)
		assert.NotEmpty(t, columns)
	})
}
