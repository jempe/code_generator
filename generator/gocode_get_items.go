package generator

// db fields to get
func (dbData *DBData) getItemsFields(structName string) (selectFields []string, returnFields []string) {
	structData := dbData.getStructData(structName)

	for _, field := range structData.Fields {
		if !field.Key {

			selectFields = append(selectFields, `if emptyOrContains(returnFields, "`+field.Name+`") {
		selectQuery += ", `+dbField(field)+`"
	}`)

			returnFields = append(returnFields, `if emptyOrContains(returnFields, "`+field.Name+`") {
		fields = append(fields, &`+goField(field)+`)
	}`)
		}
	}

	return
}
