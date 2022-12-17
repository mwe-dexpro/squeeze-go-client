package squeeze_go_client

import (
	"encoding/json"
	"fmt"
)

type DocumentClassApi struct {
	client *Client
}

func newDocumentClassApi(client *Client) *DocumentClassApi {
	return &DocumentClassApi{client: client}
}

func (api *DocumentClassApi) GetDocumentClasses() ([]*DocumentClassDto, *Error) {
	req, err := api.client.newRequest("GET", "/documentClasses")
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

	var data []*DocumentClassDto
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *DocumentClassApi) GetAllFieldGroups(documentClassId int) ([]*DocumentFieldGroup, *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/fieldGroups", documentClassId))
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

	var data []*DocumentFieldGroup
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}
func (api *DocumentClassApi) GetAllDocumentClassFields(documentClassId int) ([]*DocumentField, *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/fields", documentClassId))
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

	var data []*DocumentField
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *DocumentClassApi) GetAllDocumentClassTables(documentClassId int) ([]*DocumentTable, *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/tables", documentClassId))
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

	var data []*DocumentTable
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *DocumentClassApi) GetAllDocumentClassTableColumns(documentClassId int, tableId int) ([]*DocumentTableColumn, *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/tables/%d/columns", documentClassId, tableId))
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

	var data []*DocumentTableColumn
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}
