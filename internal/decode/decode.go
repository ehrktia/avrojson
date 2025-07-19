package decode

import (
	"fmt"

	"github.com/hamba/avro/v2"
)

func ParseSchemaData(data []byte) (avro.Schema, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data received")
	}
	return avro.ParseBytes(data)

}
