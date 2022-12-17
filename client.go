package squeeze_go_client

import (
	"fmt"
	"net/http"
)

type Client struct {
	// basePath is the base path of the Squeeze API, example: https://tenant.squeeze.one/api/v2
	basePath string

	// http is the HTTP client used to make requests
	http *http.Client

	Public *PublicApi
}

func NewClient(basePath string) *Client {
	c := &Client{basePath: basePath}

	c.http = &http.Client{}

	c.Public = newPublicApi(c)

	return c
}

func (c *Client) newRequest(method string, path string) (*http.Request, error) {
	return http.NewRequest(method, fmt.Sprintf("%s%s", c.basePath, path), nil)
}
