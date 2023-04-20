package jsonhelper

import (
	encJson "encoding/json"
	ioutil "io/ioutil"
)

func LoadOneObjFromFile[D any](d *D, filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = encJson.Unmarshal(fileContent, &d)
	if err != nil {
		return err
	}
	return nil
}
