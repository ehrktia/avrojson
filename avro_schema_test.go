package avrojson

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestDecodeSchemaToJson(t *testing.T) {
	data, err := ReadSchemaFile(filepath.Join("testdata", "basic.avro"))
	if err != nil {
		t.Fatal(err)
	}
	a := &AvroSchema{}
	if err := json.Unmarshal(data, a); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", a)
}
