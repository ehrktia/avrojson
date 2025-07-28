package avrojson

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type JsonWrapper struct {
	JS       *JSONSchema
	AS       *AvroSchema
	AvroData []byte
}

type JSONSchema struct {
	Schema      string      `json:"$schema"`
	ID          string      `json:"$id"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type"`
	Properties  []*Property `json:"properties"`
}

type Property struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
}

func NewJsonWrapper(avroData []byte) (*JsonWrapper, error) {
	js := newJsonSchema()
	as := &AvroSchema{}
	if len(avroData) < 1 {
		return nil, fmt.Errorf("empty avro data received")
	}
	if err := json.Unmarshal(avroData, as); err != nil {
		return nil, err
	}
	js.Title = as.Name
	return &JsonWrapper{
		AvroData: avroData,
		JS:       js,
		AS:       as,
	}, nil

}

func jsonAvroDataTypeMap(jsw *JsonWrapper) {
	for _, v := range jsw.JS.Properties {
		v.Type = avroJSONMap[v.Type]
	}
}

func assignProperties(js *JSONSchema, as *AvroSchema) {
	fmt.Printf("+++++++++++++++++:%d\n", len(as.Fields))
	for _, v := range as.Fields {
		fmt.Printf("name:%s\n", v.Name)
		fmt.Printf("description:%s\n", v.Doc)
		st, ok := v.Type.(string)
		var cdt customDataType
		if !ok {
			st = "interfacelist"
			cdt.value = v.Type
			cdt.dataType = "interfacelist"
			fmt.Printf("custom type:%s\n", avroJSONMap[cdt.dataType])
			fmt.Printf("custom type:%v\n", cdt.value)
		}
		fmt.Printf("type:%v\n", avroJSONMap[st])

		fmt.Println("----------------------")

	}
	// propertyList := make([]*Property, 0, len(as.Fields))
	// for _, f := range as.Fields {
	// 	p := &Property{
	// 		Name:        f.Name,
	// 		Description: f.Doc,
	// 		Type:        f.Type.(string),
	// 	}
	// 	propertyList = append(propertyList, p)
	// }
	// js.Properties = propertyList
}

const jsonSchema = `https://json-schema.org/draft/2020-12/schema`

var id = fmt.Sprintf("https://example.com/data-%d.schema.json", rand.Int())

const obj = `object`

func newJsonSchema() *JSONSchema {
	js := &JSONSchema{
		Schema: jsonSchema,
		ID:     id,
		Type:   obj,
	}
	return js
}
