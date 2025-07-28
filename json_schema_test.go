package avrojson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateJsonSchema(t *testing.T) {
	data, err := ReadSchemaFile(filepath.Join("testdata", "basic.avro"))
	if err != nil {
		t.Fatal(err)
	}
	a := &AvroSchema{}
	if err := json.Unmarshal(data, a); err != nil {
		t.Fatal(err)
	}
	ad, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	jsw, err := NewJsonWrapper(ad)
	if err != nil {
		t.Fatal(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := CreateAvroToJson(jsw, filepath.Join(cwd, "result.json")); err != nil {
		t.Fatal(err)
	}
}

func TestNestedAvro(t *testing.T) {
	data, err := ReadSchemaFile(filepath.Join("testdata", "nested.avro"))
	if err != nil {
		t.Fatal(err)
	}
	a := &AvroSchema{}
	if err := json.Unmarshal(data, a); err != nil {
		t.Fatal(err)
	}
	// t.Logf("avro raw data:%#v\n", a)
	ad, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	// t.Logf("string value:%s\n", ad)
	if err := json.Unmarshal(ad, a); err != nil {
		t.Fatal(err)
	}
	// t.Logf("fields:%#v\n", a.Fields)
	for _, v := range a.Fields {
		t.Logf("name:%s\n", v.Name)
		t.Logf("type:%s\n", v.Type)
		tv, ok := v.Type.(string)
		if !ok {
			tv = "interfacelist"
		}
		mapValue := avroJSONMap[tv]
		t.Logf("map value:%s\n", mapValue)
		t.Log("-----------------")
	}
	jsw, err := NewJsonWrapper(ad)
	if err != nil {
		t.Fatal(err)
	}
	assignProperties(jsw.JS, jsw.AS)
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if err := CreateAvroToJson(jsw, filepath.Join(cwd, "nested.json")); err != nil {
	// 	t.Fatal(err)
	// }
}
