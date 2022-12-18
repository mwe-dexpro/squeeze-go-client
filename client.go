package squeeze_go_client

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	// basePath is the base path of the Squeeze API, example: https://tenant.squeeze.one/api/v2
	basePath string

	// ApiKey is used to authenticate with Squeeze
	ApiKey string

	// http is the HTTP client used to make requests
	http *http.Client

	Document      *DocumentApi
	DocumentClass *DocumentClassApi
	Jobs          *JobsApi
	Public        *PublicApi
	Queue         *QueueApi
	ScriptingApi  *ScriptingApi
}

func NewClient(basePath string) *Client {
	c := &Client{basePath: basePath}

	c.http = &http.Client{}

	c.Document = newDocumentApi(c)
	c.DocumentClass = newDocumentClassApi(c)
	c.Jobs = newJobsApi(c)
	c.Public = newPublicApi(c)
	c.Queue = newQueueApi(c)
	c.ScriptingApi = newScriptingApi(c)

	return c
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.basePath, path), body)
	if err != nil {
		return nil, err
	}

	if c.ApiKey != "" {
		request.Header.Set("X-Api-Key", c.ApiKey)
	}

	return request, err
}
