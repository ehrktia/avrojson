### usage

```go
package main

import (
    "fmt"
    "path/filepath"

    "github.com/ehrktia/avrojson"
)

func main() {

    av, err := avrojson.SchemaFile(filepath.Join(".", "basic.avro"))
    if err != nil {
        panic(err)
    }
    jsw, err := avrojson.NewJsonWrapper(av)
    if err != nil {
        panic(err)
    }
    avrojson.TransferAvroToJson(jsw)
    fmt.Printf("json:%#v\n", jsw.JS)
    for _, v := range jsw.JS.Properties {
        fmt.Printf("js fields:%#v\n", v)
    }

}
```
