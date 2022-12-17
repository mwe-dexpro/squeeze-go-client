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

type DocumentClassDto struct {
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	MaySeeDocuments       bool   `json:"maySeeDocuments"`
	TranslationKey        string `json:"translationKey"`
	TranslatedDescription string `json:"translatedDescription"`
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

type DocumentField struct {
	Id                    int    `json:"id"`
	DocumentClassId       int    `json:"documentClassId"`
	SortOrder             int    `json:"sortOrder"`
	FieldGroupId          int    `json:"fieldGroupId"`
	LocatorId             int    `json:"locatorId"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	DataType              string `json:"dataType"`
	TranslationKey        string `json:"translationKey"`
	TranslatedDescription string `json:"translatedDescription"`
	Value                 struct {
		Id          int    `json:"id"`
		Value       string `json:"value"`
		BoundingBox struct {
			Page int `json:"page"`
			X0   int `json:"x0"`
			Y0   int `json:"y0"`
			X1   int `json:"x1"`
			Y1   int `json:"y1"`
		} `json:"boundingBox"`
		Alternative  bool     `json:"alternative"`
		FieldId      int      `json:"fieldId"`
		Confidence   int      `json:"confidence"`
		SubFieldName string   `json:"subFieldName"`
		State        string   `json:"state"`
		ErrorText    string   `json:"errorText"`
		Subfields    []string `json:"subfields"`
	} `json:"value"`
	DefaultValue            string `json:"defaultValue"`
	SameLineAsPreviousField bool   `json:"sameLineAsPreviousField"`
	Alternatives            []struct {
		Id          int    `json:"id"`
		Value       string `json:"value"`
		BoundingBox struct {
			Page int `json:"page"`
			X0   int `json:"x0"`
			Y0   int `json:"y0"`
			X1   int `json:"x1"`
			Y1   int `json:"y1"`
		} `json:"boundingBox"`
		Alternative  bool     `json:"alternative"`
		FieldId      int      `json:"fieldId"`
		Confidence   int      `json:"confidence"`
		SubFieldName string   `json:"subFieldName"`
		State        string   `json:"state"`
		ErrorText    string   `json:"errorText"`
		Subfields    []string `json:"subfields"`
	} `json:"alternatives"`
	SubFieldName    string `json:"subFieldName"`
	Mandatory       bool   `json:"mandatory"`
	Readonly        bool   `json:"readonly"`
	Hidden          bool   `json:"hidden"`
	ForceValidation bool   `json:"forceValidation"`
	State           string `json:"state"`
	ExternalName    string `json:"externalName"`
	Lookup          struct {
		Active               bool  `json:"active"`
		AllowCustomValues    bool  `json:"allowCustomValues"`
		TableId              int   `json:"tableId"`
		ResultKeyColumnId    int   `json:"resultKeyColumnId"`
		SearchColumnIds      []int `json:"searchColumnIds"`
		ResultValueColumnIds []int `json:"resultValueColumnIds"`
		LookupFieldFilters   []struct {
			Id                 int    `json:"id"`
			DocumentFieldId    int    `json:"documentFieldId"`
			MasterDataColumnId int    `json:"masterDataColumnId"`
			Operand            string `json:"operand"`
			ValueFieldId       int    `json:"valueFieldId"`
			RowBasedFilter     bool   `json:"rowBasedFilter"`
		} `json:"lookupFieldFilters"`
		MinInputLength        int  `json:"minInputLength"`
		IgnoreInputValue      bool `json:"ignoreInputValue"`
		MaxLookupResultValues int  `json:"maxLookupResultValues"`
	} `json:"lookup"`
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
