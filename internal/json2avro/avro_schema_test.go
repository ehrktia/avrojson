package json2avro

import (
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestDecodeSchemaToJson(t *testing.T) {
	data, err := SchemaFile(filepath.Join("testdata", "basic.avro"))
	if err != nil {
		t.Fatal(err)
	}
	a := &AvroSchema{}
	if err := json.Unmarshal(data, a); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", a)
}
