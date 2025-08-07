package avrojson

/*
avro type	json type	example
null	null	null
boolean	boolean	true
int,long	integer	1
float,double	number	1.1
bytes	string	"\u00FF"
string	string	"foo"
record	object	{"a": 1}
enum	string	"FOO"
array	array	[1]
map	object	{"a": 1}
fixed	string	"\u00ff"
*/
var avroJSONMap = map[string]string{
	"null":          "object",
	"boolean":       "boolean",
	"int":           "integer",
	"long":          "integer",
	"float":         "number",
	"double":        "number",
	"bytes":         "string",
	"string":        "string",
	"record":        "object",
	"enum":          "string",
	"array":         "array",
	"map":           "object",
	"fixed":         "string",
	"interfacelist": "array",
}

type customDataType struct {
	value    []byte
	dataType string
}
