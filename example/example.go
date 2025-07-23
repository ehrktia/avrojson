package main

import (
	"os"
	"path/filepath"

	"github.com/ehrktia/avrojson"
)

func main() {

	av, err := avrojson.ReadSchemaFile(filepath.Join(".", "basic.avro"))
	if err != nil {
		panic(err)
	}
	jsw, err := avrojson.NewJsonWrapper(av)
	if err != nil {
		panic(err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if err := avrojson.CreateAvroToJson(jsw, filepath.Join(cwd, "result.json")); err != nil {
		panic(err)
	}

}
