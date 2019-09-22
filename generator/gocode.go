package generator

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

type GoData struct {
	StructName       string
	StructPluralName string
	TableName        string
	TableItem        string
	KeyType          string
	KeyName          string
	KeyFieldName     string
	StructVar        string
	DbFields         []Field
	OpenBraces       string
	CloseBraces      string
}

func (dbData *DBData) ProcessTemplates(structName string, templateFiles ...string) string {
	goTemplate := template.Must(template.ParseFiles(templateFiles...))

	goStructData := GoData{OpenBraces: "{{", CloseBraces: "}}"}

	for _, structData := range dbData.StructsData {
		if structData.Name == structName {
			goStructData.StructName = structName
			goStructData.StructPluralName = structName + "s"

			// common variables for many functions
			goStructData.TableName = dbData.dbTable(structName)
			goStructData.TableItem = dbData.dbTableItem(structName)
			goStructData.KeyName = dbData.keyName(structName)
			goStructData.KeyFieldName = dbData.keyFieldName(structName)
			goStructData.KeyType = dbData.keyFieldType(structName)
			goStructData.StructVar = dbData.structVar(structName)
			goStructData.DbFields = dbData.getDbFields(structName)

		}
	}

	var tpl bytes.Buffer
	err := goTemplate.Execute(&tpl, goStructData)

	if err != nil {
		log.Fatal(err)
	}

	return tpl.String()
}

func camelCase(name string) string {
	if name == "ID" {
		return "id"
	} else {
		return strings.ToLower(name[0:1]) + name[1:len(name)]
	}
}

//keyName returns the name of primary key field
func (dbData *DBData) keyName(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return field.Name
		}
	}

	return "undefined_key_field"
}

//keyFieldName returns the db Field name of primary key field
func (dbData *DBData) keyFieldName(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return field.FieldName
		}
	}

	return "undefined_key_db_field_type"
}

//keyFieldType returns the type of primary key field
func (dbData *DBData) keyFieldType(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return field.Type
		}
	}

	return "undefined_key_db_field_type"
}

// structVar returns the name of the variable of a struct
func (dbData *DBData) structVar(structName string) string {
	return camelCase(structName)

}
