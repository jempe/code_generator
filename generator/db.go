package generator

import (
	"strings"
)

// dbTable returns the name of the table of a struct
func (dbData *DBData) dbTable(structName string) string {
	structData := dbData.getStructData(structName)
	if structData.Table != "" {
		return structData.Table
	} else {
		return "undefined_table"
	}
}

// dbTableItem returns the name of the table item of a struct
func (dbData *DBData) dbTableItem(structName string) string {
	structData := dbData.getStructData(structName)
	if structData.Item != "" {
		return structData.Item
	} else {
		return "undefined_item"
	}
}

//GetDbFields returns the list of db fields
func (dbData *DBData) GetDbFields(structName string) (fieldsList []Field) {
	structData := dbData.getStructData(structName)

	for _, field := range structData.Fields {
		if field.FieldName == "" {
			field.FieldName = strings.ToLower(field.Name)
		}

		if field.ValidateFunction == "" {
			field.ValidateFunction = "Valid" + field.Name + "Default"
		}

		if field.Type == "reference" {
			field.ReferenceType = dbData.keyFieldType(field.Reference)
		} else {
			field.ReferenceType = ""
		}

		fieldsList = append(fieldsList, field)
	}
	return
}
