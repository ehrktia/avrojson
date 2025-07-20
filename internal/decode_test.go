package internal

import (
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("empty file should give error", func(t *testing.T) {
		data, err := SchemaFile(filepath.Join("testdata", "empty.avro"))
		if err != nil {
			t.Fatal(err)
		}
		if _, err := ParseSchemaData(data); err == nil {
			t.Fatal("expected to get error got nil")
		}
	})
	t.Run("valid file should give schema", func(t *testing.T) {
		data, err := SchemaFile(filepath.Join("testdata", "basic.avro"))
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
