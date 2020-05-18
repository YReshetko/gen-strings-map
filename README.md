
## Go generate tool

#### Install
```shell script
go get -u github.com/YReshetko/gen-strings-map/cmd/constanter
```

#### Usage
The tool scans provided base path for sub paths pattern, collects all packages which match the pattern.
After that the tool extracts all var\constants with string type and generates a new go file which contains a map of links to the var\const.
Keys of the map correspond to wild cards in the pattern. The last key is a name of var\constant.
###### Example
If you have following structure:
```bash
|--root
|  |--base_path
|  |  |--type_1
|  |  |  |--v0.0.1
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v1.0.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v1.0.1
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |--type_2
|  |  |  |--v0.0.1
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v1.0.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v2.0.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |--type_3
|  |  |  |--v1.0.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v1.1.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
|  |  |  |--v1.2.0
|  |  |  |  |--file_with_constants_1.go
|  |  |  |  |--file_with_constants_2.go
```
The tool can be used as cli tool for the structure like:
```bash
constanter -baseDir=./root/base_path -pattern=/*/*
```
or like a `go:generate`
```go
//go:generate constanter -baseDir=./root/base_path -pattern=/*/*
```
If each go file contains at least one string var ot constant then output map will have following structure:
```go
package base_path

import (...)

var Constants = map[string]map[string]map[string]string{
	"type_1": {
		"v0.0.1": {
			"<constant_name>":   al_type_1_v0_0_1.<constant_name>,
			...
		},
		"v1.0.0": {
			"<constant_name>":   al_type_1_v1_0_0.<constant_name>,
			...
		},
		"v1.0.1": {
			"<constant_name>":   al_type_1_v1_0_1.<constant_name>,
			...
		},
	},
	"type_2": {
		"v0.0.1": {
			"<constant_name>":   al_type_2_v0_0_1.<constant_name>,
			...
		},
		"v1.0.0": {
			"<constant_name>":   al_type_2_v1_0_0.<constant_name>,
			...
		},
		"v2.0.0": {
			"<constant_name>":   al_type_2_v2_0_0.<constant_name>,
			...
		},
	},
	"type_3": {
		"v1.0.0": {
			"<constant_name>":   al_type_3_v1_0_0.<constant_name>,
			...
		},
		"v1.1.0": {
			"<constant_name>":   al_type_3_v1_1_0.<constant_name>,
			...
		},
		"v1.2.0": {
			"<constant_name>":   al_type_3_v1_2_0.<constant_name>,
			...
		},
	},
}
```

A.go file with the map will be generated into `./root/base_path` with default name `constanter.go`
