package squeeze_go_client

import (
	"encoding/json"
	"fmt"
)

type ScriptingApi struct {
	client *Client
}

func newScriptingApi(client *Client) *ScriptingApi {
	return &ScriptingApi{client: client}
}

func (api *ScriptingApi) GetScripts() ([]*ScriptDto, *Error) {
	req, err := api.client.newRequest("GET", "/scripting/scripts")
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

	var data []*ScriptDto
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return data, nil
}

func (api *ScriptingApi) ExecuteScript(scriptId string, async bool) *Error {
	req, err := api.client.newRequest("POST", fmt.Sprintf("/scripting/scripts/%s/execute", scriptId))
	if err != nil {
		return newErr(err)
	}

	if async {
		req.URL.Query().Set("async", "true")
	} else {
		req.URL.Query().Set("async", "false")
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
