package generator

type DBData struct {
	StructsData      []StructData
	CodeTemplatePath string
	DBType           string
}

type StructData struct {
	Name   string  `json:"name"`
	Table  string  `json:"table"`
	Item   string  `json:"item"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name      string `json:"name"`
	FieldName string `json:"field_name,omitempty"`
	Key       bool   `json:"key,omitempty"`
	Type      string `json:"type"`
	Maxlength int    `json:"maxlength,omitempty"`
	Minlength int    `json:"minlength,omitempty"`
	DBExtras  string `json:"db_extras,omitempty"`
	DBDefault string `json:"db_default,omitempty"`
	Reference string `json:"reference,omitempty"`
}

func (dbData *DBData) getStructData(structName string) StructData {
	for _, structData := range dbData.StructsData {
		if structName == structData.Name {
			return structData
		}
	}
	return StructData{}
}
