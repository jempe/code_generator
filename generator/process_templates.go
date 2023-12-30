package generator

import (
	"bytes"
	"errors"
	"text/template"
)

func (dbData *DBData) ProcessTemplates(table string, templateFiles ...string) (output string, err error) {

	goTemplate := template.Must(template.ParseFiles(templateFiles...))

	dbData.OpenBraces = "{{"
	dbData.CloseBraces = "}}"

	if table != "" {
		var found bool
		for tableIndex, tableData := range dbData.Tables {
			tableData.NaturalIndex = tableIndex + 1
			if tableData.The_items_name == table {
				dbData.SelectedTable = tableData
				found = true
			}

			for fieldIndex, fieldData := range tableData.Fields {
				fieldData.NaturalIndex = fieldIndex + 1
			}
		}

		if found == false {
			return "", errors.New("Table " + table + " not found")
		}
	}

	var tpl bytes.Buffer
	err = goTemplate.Execute(&tpl, dbData)

	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
