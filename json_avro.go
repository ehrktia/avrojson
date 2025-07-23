package avrojson

import (
	"encoding/json"
	"os"
)

// CreateAvroToJson used to create json schema file from json wrapper.
// the new json schema file is created in the fileName location supplied
func CreateAvroToJson(jsw *JsonWrapper, fileName string) error {
	assignProperties(jsw.JS, jsw.AS)
	jsonAvroDataTypeMap(jsw)
	return createJsonSchemFile(jsw, fileName)
}

func createJsonSchemFile(jsw *JsonWrapper, fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	b, err := json.Marshal(jsw.JS)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}
