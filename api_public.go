package squeeze_go_client

import (
	"fmt"
)

type PublicApi struct {
	client *Client
}

func newPublicApi(client *Client) *PublicApi {
	return &PublicApi{client: client}
}

func (api *PublicApi) GetHealth() *Error {
	req, err := api.client.newRequest("GET", "/public/system/health", nil)
	if err != nil {
		return newErr(err)
	}

	res, err := api.client.http.Do(req)
	if err != nil {
		return newApiErr(err, res)
	}

	if res.StatusCode != 204 {
		return newApiErr(fmt.Errorf("unexpected status code: %d", res.StatusCode), res)
	}

	return nil
}
