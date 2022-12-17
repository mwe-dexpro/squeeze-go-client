package squeeze_go_client

type BoundingBox struct {
	Page int `json:"page"`
	X0   int `json:"x0"`
	Y0   int `json:"y0"`
	X1   int `json:"x1"`
	Y1   int `json:"y1"`
}

type DocumentField struct {
	Id                      int                   `json:"id"`
	DocumentClassId         int                   `json:"documentClassId"`
	SortOrder               int                   `json:"sortOrder"`
	FieldGroupId            int                   `json:"fieldGroupId"`
	LocatorId               int                   `json:"locatorId"`
	Name                    string                `json:"name"`
	Description             string                `json:"description"`
	DataType                string                `json:"dataType"`
	TranslationKey          string                `json:"translationKey"`
	TranslatedDescription   string                `json:"translatedDescription"`
	Value                   *DocumentFieldValue   `json:"value"`
	DefaultValue            string                `json:"defaultValue"`
	SameLineAsPreviousField bool                  `json:"sameLineAsPreviousField"`
	Alternatives            []*DocumentFieldValue `json:"alternatives"`
	SubFieldName            string                `json:"subFieldName"`
	Mandatory               bool                  `json:"mandatory"`
	Readonly                bool                  `json:"readonly"`
	Hidden                  bool                  `json:"hidden"`
	ForceValidation         bool                  `json:"forceValidation"`
	State                   string                `json:"state"`
	ExternalName            string                `json:"externalName"`
	Lookup                  *LookupDefinition     `json:"lookup"`
}

type DocumentFieldValue struct {
	Id           int          `json:"id"`
	Value        string       `json:"value"`
	BoundingBox  *BoundingBox `json:"boundingBox"`
	Alternative  bool         `json:"alternative"`
	FieldId      int          `json:"fieldId"`
	Confidence   int          `json:"confidence"`
	SubFieldName string       `json:"subFieldName"`
	State        string       `json:"state"`
	ErrorText    string       `json:"errorText"`
	Subfields    []string     `json:"subfields"`
}

type LookupDefinition struct {
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
}

type DocumentTable struct {
	Id              int `json:"id"`
	DocumentClassId int `json:"documentClassId"`
	// removed because not actively being used
	// FieldGroupId    int    `json:"fieldGroupId"`
	LocatorId             int                    `json:"locatorId"`
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Mandatory             bool                   `json:"mandatory"`
	Readonly              bool                   `json:"readonly"`
	Hidden                bool                   `json:"hidden"`
	ForceValidation       bool                   `json:"forceValidation"`
	ExternalName          string                 `json:"externalName"`
	Columns               []*DocumentTableColumn `json:"columns"`
	Rows                  []*DocumentTableRow    `json:"rows"`
	State                 string                 `json:"state"`
	ErrorText             string                 `json:"errorText"`
	ErrorCode             int                    `json:"errorCode"`
	TableBehaviour        string                 `json:"tableBehaviour"`
	TranslationKey        string                 `json:"translationKey"`
	TranslatedDescription string                 `json:"translatedDescription"`
}

type DocumentTableColumn struct {
	Id                    int               `json:"id"`
	SortOrder             int               `json:"sortOrder"`
	TableId               int               `json:"tableId"`
	HeaderLocatorId       int               `json:"headerLocatorId"`
	ValueLocatorId        int               `json:"valueLocatorId"`
	Name                  string            `json:"name"`
	Description           string            `json:"description"`
	DataType              string            `json:"dataType"`
	Mandatory             bool              `json:"mandatory"`
	Readonly              bool              `json:"readonly"`
	Hidden                bool              `json:"hidden"`
	ForceValidation       bool              `json:"forceValidation"`
	ExternalName          string            `json:"externalName"`
	Lookup                *LookupDefinition `json:"lookup"`
	TranslationKey        string            `json:"translationKey"`
	TranslatedDescription string            `json:"translatedDescription"`
}

type DocumentTableRow struct {
	Value     *DocumentFieldValue `json:"value"`
	Cells     []string            `json:"cells"`
	State     string              `json:"state"`
	ErrorText string              `json:"errorText"`
	ErrorCode int                 `json:"errorCode"`
}

type DocumentFieldGroup struct {
	Id              int    `json:"id"`
	DocumentClassId int    `json:"documentClassId"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	// removed because these are not actively being used.
	// Type                  int    `json:"type"`
	// TableField            string `json:"tableField"`
	SortOrder             int    `json:"sortOrder"`
	TranslationKey        string `json:"translationKey"`
	TranslatedDescription string `json:"translatedDescription"`
}

type DocumentClassDto struct {
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	MaySeeDocuments       bool   `json:"maySeeDocuments"`
	TranslationKey        string `json:"translationKey"`
	TranslatedDescription string `json:"translatedDescription"`
}

type QueueStepDto struct {
	Name       string `json:"name"`
	Count      int    `json:"count"`
	ErrorCount int    `json:"errorCount"`
}
