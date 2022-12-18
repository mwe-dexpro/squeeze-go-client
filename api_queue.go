package squeeze_go_client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type QueueApi struct {
	client *Client
}

func newQueueApi(client *Client) *QueueApi {
	return &QueueApi{client: client}
}

func (api *QueueApi) GetQueueSteps() (map[string]*QueueStepDto, *Error) {
	req, err := api.client.newRequest("GET", "/queue/steps", nil)
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
	req, err := api.client.newRequest("GET", fmt.Sprintf("/queue/steps/%s", stepName), nil)
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

// GetQueueStepEntriesSimple is a simple implementation of a GET /queue/steps/{stepName}/entries request with support only for the most common parameters.
func (api *QueueApi) GetQueueStepEntriesSimple(stepName string, status string, documentClassId int, batchClassId int, page int, pageSize int) (*PaginationResponse[*QueueEntryDto], *Error) {
	query := url.Values{}
	if status != "" {
		query.Add("status", status)
	}
	if documentClassId != 0 {
		query.Add("documentClassId", fmt.Sprintf("%d", documentClassId))
	}
	if batchClassId != 0 {
		query.Add("batchClassId", fmt.Sprintf("%d", batchClassId))
	}
	if page != 0 {
		query.Add("page", fmt.Sprintf("%d", page))
	}
	if pageSize != 0 {
		query.Add("pageSize", fmt.Sprintf("%d", pageSize))
	}

	path := fmt.Sprintf("/queue/steps/%s/entries", stepName)
	queryStr := query.Encode()
	if queryStr != "" {
		path = path + "?" + queryStr
	}

	req, err := api.client.newRequest("GET", path, nil)
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

	var data PaginationResponse[*QueueEntryDto]
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return &data, nil
}

func (api *QueueApi) RequeueDocument(documentId int, documentClassId int, newStepName string) *Error {
	query := url.Values{}
	if documentClassId != 0 {
		query.Add("documentClassId", strconv.Itoa(documentClassId))
	}
	query.Add("newStep", newStepName)

	path := fmt.Sprintf("/documents/%d/requeue?%s", documentId, query.Encode())

	req, err := api.client.newRequest("PUT", path, nil)
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
