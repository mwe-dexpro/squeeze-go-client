package squeeze_go_client

import (
	"fmt"
	"net/http"
)

type Client struct {
	// basePath is the base path of the Squeeze API, example: https://tenant.squeeze.one/api/v2
	basePath string

	// ApiKey is used to authenticate with Squeeze
	ApiKey string

	// http is the HTTP client used to make requests
	http *http.Client

	DocumentClass *DocumentClassApi
	Public        *PublicApi
	Queue         *QueueApi
}

func NewClient(basePath string) *Client {
	c := &Client{basePath: basePath}

	c.http = &http.Client{}

	c.DocumentClass = newDocumentClassApi(c)
	c.Public = newPublicApi(c)
	c.Queue = newQueueApi(c)

	return c
}

func (c *Client) newRequest(method string, path string) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.basePath, path), nil)
	if err != nil {
		return nil, err
	}

	if c.ApiKey != "" {
		request.Header.Set("X-Api-Key", c.ApiKey)
	}

	return request, err
}
