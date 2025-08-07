package avrojson

import (
	"bytes"
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
	ad, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(ad, a); err != nil {
		t.Fatal(err)
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

func TestJsonUnMarshal(t *testing.T) {
	data := `{"items":{"name":"AvrosampleNetCore.SubAccounts","type":"record","fields":[{"name":"AccountId","type":"int"},{"name":"AccountType","type":["null","string"]}]},"type":"array"}`
	db := []byte(data)
	var alv *ArrayList
	if err := json.Unmarshal(db, &alv); err != nil {
		t.Fatal(err)
	}
	result, err := json.Marshal(alv)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(db, result) {
		t.Fatalf("expected:%s\ngot:%s\n", data, result)
	}
}
