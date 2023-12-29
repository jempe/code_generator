package generator

import (
	"encoding/json"
	"io/ioutil"
)

func ReadFile(filePath string) (dbData DBData, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	err = json.Unmarshal(content, &dbData)
	if err != nil {
		return
	}

	return dbData, nil
}
