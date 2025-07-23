package avrojson

import (
	"fmt"
	"io"
	"os"
)

// ReadSchemaFile reads avro schema file from a location and emit it bytes along
// with error in case when encountered
func ReadSchemaFile(location string) ([]byte, error) {
	if location == "" {
		return nil, fmt.Errorf("invalid/empty file location received")
	}
	// open file and read
	f, err := os.Open(location)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
