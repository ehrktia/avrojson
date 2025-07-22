package avrojson

func TransferAvroToJson(avroData []byte) error {
	jsw, err := newJsonWrapper(avroData)
	if err != nil {
		return err
	}
	assignProperties(jsw.JS, jsw.AS)
	jsonAvroDataTypeMap(jsw)
	return nil
}
