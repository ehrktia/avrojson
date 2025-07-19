package decode

import (
	"avrojson/internal/read"
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {

	t.Run("empty file should give error", func(t *testing.T) {
		// TODO: make it dynamic and generic for all internal packages
		// to help with testing
		data, err := read.SchemaFile(filepath.Join("..", "/", "read", "testdata", "empty.avro"))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := ParseSchemaData(data); err == nil {
			t.Fatal("expected to get error got nil")
		}
	})
	t.Run("valid file should give schema", func(t *testing.T) {
		// TODO: make it dynamic and generic for all internal packages
		// to help with testing
		data, err := read.SchemaFile(filepath.Join("..", "/", "read", "testdata", "basic.avro"))
		if err != nil {
			t.Fatal(err)
		}
		sData, err := ParseSchemaData(data)
		if err != nil {
			t.Fatalf("err:%v\n", err)
		}
		t.Logf("data:%s\n", sData.String())
	})
}
