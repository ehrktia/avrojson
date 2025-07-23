package avrojson

import (
	"fmt"
	"io"
	"os"
)

// read a schema file from location and emit it bytes and error

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
