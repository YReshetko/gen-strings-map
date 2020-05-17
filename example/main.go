package main

import (
	"fmt"
	"github.com/YReshetko/map-generator/example/test"
)

//go:generate ../bin/constanter -baseDir=./test -pattern=/schema/*/*/schema
func main()  {
	fmt.Println(test.Constants["order"]["0.1.0"]["Payload"])
}
