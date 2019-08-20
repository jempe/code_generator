package generator

import (
	"bytes"
	"strings"
	"text/template"
)

type GoData struct {
	StructName       string
	StructPluralName string
	TableName        string
	KeyType          string
	KeyGoType        string
	KeyDbName        string
	KeyName          string
	ParameterKeyName string
	StructVar        string
	DbFields         []SQLField
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
			goStructData.KeyDbName = dbData.keyDbField(structName)
			goStructData.KeyName = dbData.keyGoFieldName(structName)
			goStructData.KeyType = dbData.keyFieldType(structName)
			goStructData.KeyGoType = dbData.keyGoFieldType(structName)
			goStructData.ParameterKeyName = dbData.keyGoParameterFieldName(structName)
			goStructData.StructVar = dbData.structVar(structName)
			goStructData.DbFields = dbData.getDbFields(structName)

		}
	}

	var tpl bytes.Buffer
	goTemplate.Execute(&tpl, goStructData)

	return tpl.String()
}

func camelCase(name string) string {
	if name == "ID" {
		return "id"
	} else {
		return strings.ToLower(name[0:1]) + name[1:len(name)]
	}
}

// goField returns the name of variable of a Field
func goField(field Field) string {
	return camelCase(field.Name)
}

// goFieldType returns the type of variable of a Field
func (dbData *DBData) goFieldType(field Field) string {
	fieldType := "string"

	if field.Type == "timestamp" || field.Type == "timestamp_now" {
		fieldType = "time.Time"

	} else if field.Type == "latitude" || field.Type == "longitude" || field.Type == "price" || field.Type == "bigint" {
		fieldType = "int"
	} else if field.Type == "bool" {
		fieldType = "bool"
	} else if field.Type == "reference" {
		fieldType = dbData.keyGoFieldType(field.Reference)
	}

	return fieldType
}

//keyGoParameterFieldName returns the go var name of the key field for the function parameters
func (dbData *DBData) keyGoParameterFieldName(structName string) string {
	return strings.ToLower(structName) + dbData.keyGoFieldName(structName)
}

//keyGoField returns the go var name of primary key field
func (dbData *DBData) keyGoField(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return goField(field)
		}
	}

	return "undefined_key_db_field_type"
}

//keyGoFieldName returns the go var name (uppercase) of primary key field
func (dbData *DBData) keyGoFieldName(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return field.Name
		}
	}

	return "undefined_key_db_field_type"
}

//keyGoFieldType returns the go type of primary key field
func (dbData *DBData) keyGoFieldType(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return dbData.goFieldType(field)
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

func (dbData *DBData) hashField(field Field, structName string, returnFields string) string {
	code := `	hashed` + field.Name + `, err := bcrypt.GenerateFromPassword([]byte(` + dbData.structVar(structName) + `.` + field.Name + `), 8)
	if err != nil {
		log.Println(validationErrorPrefix, err)
		return` + returnFields + `
	}
	`

	if dbData.DBType == "boltdb" {
		code += dbData.structVar(structName) + "." + field.Name
	} else {
		code += goField(field)
	}

	code += `= string(hashed` + field.Name + `)` + "\n\n"

	return code
}
