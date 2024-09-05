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
	req, err := api.client.newRequest("GET", "/documentClasses", nil)
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
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/fieldGroups", documentClassId), nil)
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
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/fields", documentClassId), nil)
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
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/tables", documentClassId), nil)
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
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/tables/%d/columns", documentClassId, tableId), nil)
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

func (api *DocumentClassApi) GetExportInterfaceStates(documentClassId int, exportInterfaceId int, status string, ignoreDeleted bool, page int, pageSize int) (*PaginationResponse[ExportStatus], *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/exportInterfaces/%d/states", documentClassId, exportInterfaceId), nil)
	if err != nil {
		return nil, newErr(err)
	}

	// Set default value for status if it is empty
	if status == "" {
		status = "ASYNC_WAITING"
	}

	// Get the query parameters
	query := req.URL.Query()

	// Set the query parameters
	query.Set("status", status)
	if ignoreDeleted {
		query.Set("async", "true")
	} else {
		query.Set("async", "false")
	}
	query.Set("page", fmt.Sprintf("%d", page))
	query.Set("pageSize", fmt.Sprintf("%d", pageSize))

	// Encode the modified query parameters back to the URL
	req.URL.RawQuery = query.Encode()

	res, err := api.client.http.Do(req)
	if err != nil {
		return nil, newApiErr(err, res)
	}

	if res.StatusCode != 200 {
		return nil, newApiErr(fmt.Errorf("unexpected status code: %d", res.StatusCode), res)
	}

	var data *PaginationResponse[ExportStatus]
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *DocumentClassApi) GetDocumentsInExportConfig(documentClassId int, exportInterfaceId int, status string, ignoreDeleted bool, page int, pageSize int) (*PaginationResponse[DocumentRefDto], *Error) {
	req, err := api.client.newRequest("GET", fmt.Sprintf("/documentClasses/%d/exportInterfaces/%d/documents", documentClassId, exportInterfaceId), nil)
	if err != nil {
		return nil, newErr(err)
	}

	// Set default value for status if it is empty
	if status == "" {
		status = "ASYNC_WAITING"
	}

	// Get the query parameters
	query := req.URL.Query()

	// Set the query parameters
	query.Set("status", status)
	if ignoreDeleted {
		query.Set("async", "true")
	} else {
		query.Set("async", "false")
	}
	query.Set("page", fmt.Sprintf("%d", page))
	query.Set("pageSize", fmt.Sprintf("%d", pageSize))

	// Encode the modified query parameters back to the URL
	req.URL.RawQuery = query.Encode()

	res, err := api.client.http.Do(req)
	if err != nil {
		return nil, newApiErr(err, res)
	}

	if res.StatusCode != 200 {
		return nil, newApiErr(fmt.Errorf("unexpected status code: %d", res.StatusCode), res)
	}

	var data *PaginationResponse[DocumentRefDto]
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}
