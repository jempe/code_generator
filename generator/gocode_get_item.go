package generator

// validator of ID for get Item function
func (dbData *DBData) getItemIDValidator(structName string) string {
	fieldType := dbData.keyFieldType(structName)

	if fieldType == "uuid" {
		return "validID, err := validator.UUID(" + dbData.keyGoParameterFieldName(structName) + ")\n\tif !validID {\n\t\treturn " + dbData.structVar(structName) + ", err\n\t}"
	}

	return ""
}
