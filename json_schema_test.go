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
	ad, err := json.Marshal(a)
	if err != nil {
		t.Fatal(err)
	}
	jsw, err := NewJsonWrapper(ad)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", jsw.JS)
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if err := CreateAvroToJson(jsw, filepath.Join(cwd, "nested.json")); err != nil {
	// 	t.Fatal(err)
	// }
}
