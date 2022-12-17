package squeeze_go_client

import (
	"encoding/json"
	"fmt"
)

type QueueApi struct {
	client *Client
}

func newQueueApi(client *Client) *QueueApi {
	return &QueueApi{client: client}
}

func (api *QueueApi) GetQueueSteps() (map[string]*QueueStepDto, *Error) {
	req, err := api.client.newRequest("GET", "/queue/steps")
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

	var data map[string]*QueueStepDto
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *QueueApi) GetQueueStep(stepName string) (*QueueStepDto, *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/queue/steps/%s", stepName))
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

	var data QueueStepDto
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return &data, nil
}
