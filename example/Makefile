.PHONY: all
all: build_tool gen build

.PHONY: build_tool
build_tool:
	go build -o ../bin/constanter ../cmd/constanter

.PHONY: gen
gen:
	go generate ./...

.PHONY: build
build:
	go run .