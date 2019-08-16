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

	var structsData []StructData

	err = json.Unmarshal(content, &structsData)
	if err != nil {
		return
	}

	dbData = DBData{StructsData: structsData}

	return
}
