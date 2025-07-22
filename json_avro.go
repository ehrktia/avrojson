package avrojson

func TransferAvroToJson(jsw *JsonWrapper) {
	assignProperties(jsw.JS, jsw.AS)
	jsonAvroDataTypeMap(jsw)
}
