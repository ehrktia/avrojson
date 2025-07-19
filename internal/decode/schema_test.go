package decode

import (
	"avrojson/internal/read"
	"encoding/json"
	"path/filepath"
	"testing"
)

func TestDecodeSchemaToJson(t *testing.T) {
	data, err := read.SchemaFile(filepath.Join("..", "/", "read", "testdata", "basic.avro"))
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
