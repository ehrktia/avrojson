package avrojson

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// JsonWrapper is a wrapper which holds
// skeleton for Json Schema and Avro Schema
// data read from Avro schema provided for conversion
type JsonWrapper struct {
	JS       *JSONSchema
	AS       *AvroSchema
	AvroData []byte
}

// JSONSchema holds all required field to create
// json schema
type JSONSchema struct {
	Schema      string      `json:"$schema"`
	ID          string      `json:"$id"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	Type        string      `json:"type"`
	Properties  []*Property `json:"properties"`
}

// Property holds the fields used in json schema
type Property struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type"`
}

// NewJsonWrapper creates a json wrapper using the data from
// provided avro file
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
	for _, v := range as.Fields {
		fmt.Printf("name:%s\n", v.Name)
		fmt.Printf("description:%s\n", v.Doc)
		st, ok := v.Type.(string)
		var cdt customDataType
		if !ok {
			st = "interfacelist"
			db, err := json.Marshal(v.Type)
			if err != nil {
				fmt.Printf("unmarshal type error:%v\n", err)
			}
			cdt.value = db
			cdt.dataType = "interfacelist"
			// fmt.Printf("custom type:%s\n", avroJSONMap[cdt.dataType])
			// ar := make([]*Property, 0)
			jsonArray := []json.RawMessage{}
			if err := json.Unmarshal(cdt.value, &jsonArray); err != nil {
				fmt.Printf("unmarshal err:%v\n", err)
			}
			// fmt.Printf("len:%d\n", len(jsonArray))
			// for _, v := range jsonArray {
			// 	fmt.Printf("json array:%s\n", v)
			// }
			fmt.Printf("array value:%s\n", jsonArray[1])
			// avData, err := json.Marshal(&jsonArray)
			// if err != nil {
			// 	fmt.Printf("err:%v\n", err)
			// }
			// av := &AvroArrayFields{}
			// if err := json.Unmarshal(cdt.value, av); err != nil {
			// 	fmt.Printf("avro err:%v\n", err)
			// }
			// av := &AvroArrayFields{}
			// avm := map[string]any{}
			var avm *map[string]any
			if err := json.Unmarshal(jsonArray[1], &avm); err != nil {
				fmt.Printf("unmarshal:%v\n", err)
			}
			for k, v := range *avm {
				fmt.Println("----------------")
				fmt.Printf("key:%s\n", k)
				fmt.Printf("value:%v\n", v)
				fmt.Println("----------------")

			}

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
