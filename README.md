This package is zero dependency avro schema to json schema. 

it exposes `ReadSchemaFile` to consume the avro schema.

`NewJsonWrapper` is a holder which manages the avro data bytes and new json

schema. `CreateAvroToJson` is the final call to populate avro data in to the json schema file to a location provided.

### usage

```go
package main

import (
    "fmt"
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
    cwd,err:=os.Getwd()
        if err!=nil {
            panic(err)
    }
    if err:=avrojson.CreateAvroToJson(jsw,filepath.Join(cwd,"result.json"));
    err!=nil{
        panic(err)
    }

}
```

output -pretty formatted

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://example.com/data-6182600211311189528.schema.json",
  "title": "thecodebuzz_schema",
  "type": "object",
  "properties": [
    {
      "name": "username",
      "description": "Name of the user account on Thecodebuzz.com",
      "type": "string"
    },
    {
      "name": "email",
      "description": "The email of the user logging message on the blog",
      "type": "string"
    },
    {
      "name": "timestamp",
      "description": "time in seconds",
      "type": "integer"
    }
  ]
}
```
