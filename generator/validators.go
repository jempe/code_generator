package generator

import (
	"strconv"
)

func goValidatorCode(field Field, structName string, returnFields string) string {
	var code string

	if field.Minlength > 0 {
		code += `validMinLength` + field.Name + `, err := validator.MinLength(` + structName + `.` + field.Name + `, ` + strconv.Itoa(field.Minlength) + `)
	if !validMinLength` + field.Name + ` { 
		return` + returnFields + `
	}` + "\n\n"
	}

	if field.Maxlength > 0 {
		code += `validMaxLength` + field.Name + `, err := validator.MaxLength(` + structName + `.` + field.Name + `, ` + strconv.Itoa(field.Maxlength) + `)
	if !validMaxLength` + field.Name + ` {
		return` + returnFields + `
	}` + "\n\n"
	}

	if field.Type == "email" {
		code += `validEmail` + field.Name + `, err := validator.Email(` + structName + `.` + field.Name + `)
	if !validEmail` + field.Name + ` {
		return` + returnFields + `
	}` + "\n\n"
	}

	return code
}
