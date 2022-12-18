package squeeze_go_client

import (
	"encoding/json"
	"fmt"
)

type JobsApi struct {
	client *Client
}

func newJobsApi(client *Client) *JobsApi {
	return &JobsApi{client: client}
}

func (api *JobsApi) GetAllJobs() ([]*Job, *Error) {
	req, err := api.client.newRequest("GET", "/jobs", nil)
	if err != nil {
		return nil, newErr(err)
	}

	res, err := api.client.http.Do(req)
	if err != nil {
		return nil, newApiErr(err, res)
	}

	if res.StatusCode != 200 {
		return nil, newApiErr(fmt.Errorf("unexpected status code: %d", res.StatusCode), res)
	}

	var data []*Job
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}
