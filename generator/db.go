package generator

import (
	"bytes"
	"strconv"
	"strings"
	"text/template"
)

type SQLDB struct {
	Migrations int
	Tables     []SQLTable
}

type SQLTable struct {
	Name   string
	Fields []SQLField
}

type SQLField struct {
	Name         string
	GoStructName string
	GoVarName    string
	GoVarType    string
	FieldType    string
	Type         string
	Extra        string
	Default      string
	Key          bool
}

func (dbData *DBData) SqlCode(currentMigration int) string {
	paths := []string{dbData.CodeTemplatePath + "/" + dbData.DBType + "/db.sql"}
	sqlTemplate := template.Must(template.ParseFiles(paths...))

	var sqlTables []SQLTable

	for _, structData := range dbData.StructsData {
		sqlFields := dbData.getDbFields(structData.Name)

		sqlTables = append(sqlTables, SQLTable{Name: dbData.dbTable(structData.Name), Fields: sqlFields})
	}

	sqlDB := SQLDB{Migrations: currentMigration, Tables: sqlTables}

	var tpl bytes.Buffer
	sqlTemplate.Execute(&tpl, sqlDB)

	return tpl.String()
}

// dbTable returns the name of the table of a struct
func (dbData *DBData) dbTable(structName string) string {
	structData := dbData.getStructData(structName)
	if structData.Table != "" {
		return structData.Table
	} else {
		return "undefined_table"
	}
}

// dbField returns the name of the DB table field
func dbField(field Field) string {
	if field.FieldName != "" {
		return field.FieldName
	} else {
		return strings.ToLower(field.Name)
	}

	return "undefined_db_field"
}

// dbFieldType returns the type of the DB table field
func dbFieldType(field Field) string {
	dbFieldLength := "100"

	if field.Maxlength > 0 {
		dbFieldLength = strconv.Itoa(field.Maxlength)
	}

	dbFieldType := "VARCHAR(" + dbFieldLength + ")"

	if field.Type == "uuid" {
		dbFieldType = "CHAR(36)"
	} else if field.Type == "currency" {
		dbFieldType = "CHAR(3)"
	} else if field.Type == "bool" {
		dbFieldType = "BOOLEAN"
	} else if field.Type == "timestamp" || field.Type == "timestamp_now" {
		dbFieldType = "TIMESTAMP"
	}

	return dbFieldType
}

//getDbFields returns the list of db fields
func (dbData *DBData) getDbFields(structName string) []SQLField {
	structData := dbData.getStructData(structName)
	var sqlFields []SQLField

	for _, field := range structData.Fields {
		var sqlField SQLField

		sqlField.Key = field.Key

		if field.Type == "string" || field.Type == "password" || field.Type == "email" || field.Type == "text" || field.Type == "currency" {
			sqlField.Default = " DEFAULT '" + field.DBDefault + "'"
		} else if field.Type == "bool" {
			sqlField.Default = " DEFAULT 'FALSE'"

			if strings.ToLower(field.DBDefault) == "true" {
				sqlField.Default = " DEFAULT 'TRUE'"
			}
		} else if field.Type == "timestamp" || field.Type == "timestamp_now" {
			sqlField.Extra = " NULL"

			if field.Type == "timestamp_now" {
				sqlField.Extra = " NOT NULL"
				sqlField.Default = " DEFAULT NOW()"
			}
		}

		sqlField.FieldType = dbFieldType(field)

		if field.Type == "reference" {
			sqlField.FieldType = dbData.keyDbFieldType(field.Reference)

			sqlField.Extra = " REFERENCES " + dbData.dbTable(field.Reference) + "(" + dbData.keyDbField(field.Reference) + ")\n\t\tON UPDATE CASCADE\n\t\tON DELETE CASCADE"
		}

		if field.Key {
			sqlField.Extra = " PRIMARY KEY"
			sqlField.Default = ""
		} else if field.DBExtras != "" {
			sqlField.Extra = " " + field.DBExtras
		}

		sqlField.GoStructName = field.Name
		sqlField.Name = dbField(field)
		sqlField.GoVarName = goField(field)
		sqlField.GoVarType = dbData.goFieldType(field)
		sqlField.Type = field.Type

		sqlFields = append(sqlFields, sqlField)
	}

	return sqlFields
}

//keyDbField returns the name of the db primary key field
func (dbData *DBData) keyDbField(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return dbField(field)
		}
	}

	return "undefined_key_db_field"
}

//keyDbFieldType returns the name of the db primary key field
func (dbData *DBData) keyDbFieldType(structName string) string {
	structData := dbData.getStructData(structName)
	for _, field := range structData.Fields {
		if field.Key {
			return dbFieldType(field)
		}
	}

	return "undefined_key_db_field_type"
}
