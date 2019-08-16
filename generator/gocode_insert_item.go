package generator

// validator of ID for get Item function
func (dbData *DBData) insertItemValidators(structName string) []string {
	structData := dbData.getStructData(structName)

	var fieldValidators []string

	for _, field := range structData.Fields {
		if !field.Key {
			fieldValidators = append(fieldValidators, goValidatorCode(field, dbData.structVar(structName), ""))
		}

	}

	return fieldValidators
}

// code to process fields
func (dbData *DBData) insertItemProcessFields(structName string) []string {
	structData := dbData.getStructData(structName)

	var processFieldsCode []string

	for _, field := range structData.Fields {
		if field.Key && field.Type == "uuid" {
			processFieldsCode = append(processFieldsCode, dbData.keyDbField(structName)+`, err := uuid.NewRandom()

	if err != nil {
		log.Println(validationErrorPrefix, err)
		return
	}`)

			if dbData.DBType == "boltdb" {
				processFieldsCode = append(processFieldsCode, dbData.keyGoParameterFieldName(structName)+" = "+dbData.keyDbField(structName)+".String()")
				processFieldsCode = append(processFieldsCode, dbData.structVar(structName)+"."+dbData.keyGoFieldName(structName)+" = "+dbData.keyGoParameterFieldName(structName))
			}

		}

		if field.Type != "timestamp_now" {
			if dbData.DBType != "boltdb" {
				if !field.Key {
					processFieldsCode = append(processFieldsCode, goField(field)+" := "+dbData.structVar(structName)+"."+field.Name)
				}
			}

			if field.Type == "password" {
				processFieldsCode = append(processFieldsCode, dbData.hashField(field, structName, ""))
			}

		}

	}

	return processFieldsCode
}

// db fields to insert
func (dbData *DBData) insertItemFields(structName string) (fields map[int]string) {
	structData := dbData.getStructData(structName)

	fields = make(map[int]string)

	fieldIndex := 0

	for _, field := range structData.Fields {
		if field.Type != "timestamp_now" {
			fieldIndex++

			fields[fieldIndex] = dbField(field)
		}

	}

	return fields
}

// db fields to insert
func (dbData *DBData) insertItemGoFields(structName string) (fields map[int]string) {
	structData := dbData.getStructData(structName)

	fields = make(map[int]string)

	fieldIndex := 0

	for _, field := range structData.Fields {
		if field.Type != "timestamp_now" {
			fieldIndex++

			fields[fieldIndex] = goField(field)
		}

	}

	return fields
}
