package main

import (
	"github.com/YReshetko/gen-strings-map/use_case/jsonschema/payloads"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"log"
)

const (
	validOrder = "./fixtures/Order_valid_V_0_1.json"
	invalidOrder = "./fixtures/Order_invalid_V_0_1.json"
)

//go:generate ../bin/constanter -baseDir=./jsonschema/payloads -pattern=/*/*
func main() {
	//check valid payload
	d := string(getData(validOrder))
	ok := Valid(d, payloads.Constants["order"]["0.1"]["SchemaPayloadPayload"])
	if !ok {
		log.Fatal("Unable to validate correct payload")
	}

	//check invalid payload
	d = string(getData(invalidOrder))
	ok = Valid(d, payloads.Constants["order"]["0.1"]["SchemaPayloadPayload"])
	if ok {
		log.Fatal("Unable to validate incorrect payload")
	}
	log.Println("Everything is OK")
}

func Valid(data, schema string) bool {
	s, err := gojsonschema.NewSchema(gojsonschema.NewStringLoader(schema))
	if err != nil {
		log.Fatal(err)
	}

	bl := gojsonschema.NewBytesLoader([]byte(data))

	r, err := s.Validate(bl)
	if err != nil{
		log.Fatal(err)
	}
	return r.Valid()
}

func getData(fileName string) []byte {
	p, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return p
}