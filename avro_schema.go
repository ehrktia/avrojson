package avrojson

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
