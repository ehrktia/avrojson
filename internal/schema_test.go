package internal

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
	s, err := ParseSchemaData(data)
	if err != nil {
		t.Fatal(err)
	}
	a := &AvroSchema{}
	if err := json.Unmarshal([]byte(s.String()), a); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v\n", a)

}
