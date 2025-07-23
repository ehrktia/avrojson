package avrojson

import (
	"path/filepath"
	"testing"
)

func TestSchemaFile(t *testing.T) {
	t.Run("empty file location should return error", func(t *testing.T) {
		if _, err := ReadSchemaFile(""); err == nil {
			t.Fatal("expected error got nil")
		}

	})
	t.Run("should be able to open and read file successfully",
		func(t *testing.T) {
			data, err := ReadSchemaFile(filepath.Join("testdata", "basic.avro"))
			if err != nil {
				t.Fatal(err)
			}
			if len(data) == 0 {
				t.Fatal("empty file received")
			}
		})
}
