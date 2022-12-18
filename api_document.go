package squeeze_go_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
)

type DocumentApi struct {
	client *Client
}

func newDocumentApi(client *Client) *DocumentApi {
	return &DocumentApi{client: client}
}

func (api *DocumentApi) ProcessDocument(batchClassId, documentClassId int, externalId string, fields map[string]string, file io.Reader, filename string) (*CreatedIntDto, *Error) {
	body := bytes.NewBuffer([]byte{})
	writer := multipart.NewWriter(body)

	// request params
	{
		if err := writer.WriteField("batchClassId", strconv.Itoa(batchClassId)); err != nil {
			return nil, newErr(fmt.Errorf("writing batchClassId failed: %s", err))
		}
		if documentClassId > 0 {
			if err := writer.WriteField("documentClassId", strconv.Itoa(documentClassId)); err != nil {
				return nil, newErr(fmt.Errorf("writing documentClassId failed: %s", err))
			}
		}
		if externalId != "" {
			if err := writer.WriteField("externalId", externalId); err != nil {
				return nil, newErr(fmt.Errorf("writing externalId failed: %s", err))
			}
		}

		if fields != nil {
			fieldsJson, err := json.Marshal(fields)
			if err != nil {
				return nil, newErr(fmt.Errorf("marshaling fields failed: %s", err))
			}
			if err := writer.WriteField("fields", string(fieldsJson)); err != nil {
				return nil, newErr(fmt.Errorf("writing fields failed: %s", err))
			}
		}
	}

	// file
	{
		formFile, err := writer.CreateFormFile("file", filename)
		if err != nil {
			return nil, newErr(fmt.Errorf("creating form file failed: %s", err))
		}
		_, err = io.Copy(formFile, file)
		if err != nil {
			return nil, newErr(fmt.Errorf("copying file failed: %s", err))
		}
		err = writer.Close()
		if err != nil {
			return nil, newErr(fmt.Errorf("closing writer failed: %s", err))
		}
	}

	req, err := api.client.newRequest("POST", "/documents", body)
	if err != nil {
		return nil, newErr(err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := api.client.http.Do(req)
	if err != nil {
		return nil, newApiErr(err, res)
	}

	if res.StatusCode != 200 {
		return nil, newApiErr(fmt.Errorf("unexpected status code: %d", res.StatusCode), res)
	}

	var data CreatedIntDto
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, newApiErr(fmt.Errorf("unmarshaling json failed: %s", err), res)
	}

	return &data, nil
}
