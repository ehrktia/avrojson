package internal

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	got := newJsonSchema()
	dbytes, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(dbytes),
		jsonSchema) {
		t.Fatal("expected to get valid schema")
	}
}

func TestJsonWrapper(t *testing.T) {
	t.Run("valid schema should produce a json schema with name title and description", func(t *testing.T) {
		avData, err := SchemaFile(filepath.Join("testdata", "basic.avro"))
		if err != nil {
			t.Fatal(err)
		}
		got, err := newJsonWrapper(avData)
		if err != nil {
			t.Fatal(err)
		}
		if len(got.AvroData) < 1 {
			t.Fatal("got empty avro data")
		}
		if got.JS == nil {
			t.Fatal("received nil json schema")
		}
		if got.AS == nil {
			t.Fatal("received nil avro schema")
		}
		if got.JS.Title == "" {
			t.Logf("jsw:%#v\n", got.JS)
			t.Logf("jsw:%#v\n", got.AS)
			t.Logf("as:%s\n", got.AS.Doc)
			t.Fatal("expected valid description missing")
		}

	})
	t.Run("should return err when avro data is empty", func(t *testing.T) {
		if _, err := newJsonWrapper([]byte{}); err == nil {
			t.Fatal("expected to get error")
		}
	})

}

func TestJsonProperties(t *testing.T) {
	avData, err := SchemaFile(filepath.Join("testdata", "basic.avro"))
	if err != nil {
		t.Fatal(err)
	}

	jsw, err := newJsonWrapper(avData)
	if err != nil {
		t.Fatal(err)
	}
	propertyList := make([]*Property, 0, len(jsw.AS.Fields))
	for _, f := range jsw.AS.Fields {
		p := &Property{
			Name:        f.Name,
			Description: f.Doc,
			Type:        f.Type.(string),
		}
		propertyList = append(propertyList, p)

	}
	jsw.JS.Properties = propertyList
	db, err := json.Marshal(jsw.JS)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", db)
}
