package generator

// code to process fields
func (dbData *DBData) updateItemProcessFields(structName string) []string {
	structData := dbData.getStructData(structName)

	var processFields []string

	for _, field := range structData.Fields {
		if !field.Key && field.Type != "timestamp_now" {
			fieldToUpdate := dbData.structVar(structName) + "." + field.Name

			var fieldExtraCode string

			if field.Type == "password" {

				fieldExtraCode = "\tvar " + goField(field) + " string\n\t" + dbData.hashField(field, structName, " rowsAffected, err")
				fieldToUpdate = goField(field)
			}

			if dbData.DBType == "boltdb" {
				processFields = append(processFields, `if emptyOrContains(fields, "`+goField(field)+`") {
		`+goValidatorCode(field, dbData.structVar(structName), " rowsAffected, err")+`

			`+dbData.structVar(structName)+`Data.`+field.Name+` = `+dbData.structVar(structName)+`.`+field.Name+`

	}`)
			} else {

				processFields = append(processFields, `if emptyOrContains(fields, "`+goField(field)+`") {
		`+goValidatorCode(field, dbData.structVar(structName), " rowsAffected, err")+`

		if len(fieldsToUpdate) > 0 {
			updateQuery += " ,"
		}
`+fieldExtraCode+`
		fieldsToUpdate = append(fieldsToUpdate, `+fieldToUpdate+`)

		updateQuery += " `+dbField(field)+`=$" + strconv.Itoa(len(fieldsToUpdate))
	}`)
			}
		}

	}

	return processFields
}
