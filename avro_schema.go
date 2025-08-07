package avrojson

// AvroSchema holds information from avro schema
type AvroSchema struct {
	Name        string        `json:"name,omitempty"`
	Type        any           `json:"type"`
	Items       any           `json:"items,omitempty"`
	Values      any           `json:"values,omitempty"`
	Fields      []*AvroSchema `json:"fields,omitempty"`
	Namespace   string        `json:"namespace,omitempty"`
	Doc         string        `json:"doc,omitempty"`
	Aliases     []string      `json:"aliases,omitempty"`
	Default     any           `json:"default,omitempty"`
	LogicalType string        `json:"logicalType,omitempty"`
}

type Field struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Type        any    `json:"type"`
}

type Items struct {
	Name        string  `json:"name"`
	Type        any     `json:"type"`
	Description string  `json:"description,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}

type ArrayList struct {
	ArrayItems Items `json:"items"`
	Type       any   `json:"type"`
}
