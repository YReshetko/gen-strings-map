PAYLOADS = $(shell cd ./swagger/payloads && find | grep -oP ^.*/[0-9]+\.[0-9]+$)

.PHONY: all
all: build_tool gen_schemas gen build

.PHONY: gen_schemas
gen_schemas:
	@mkdir -p ./jsonschema
	@mkdir -p ./jsonschema/payloads
	@openapi2jsonschema -block_additional_properties -go_constants -allow_null_values -spec swagger/envelope.yaml -out=./jsonschema
	for dir in $(PAYLOADS) ; do \
		mkdir -p -v ./jsonschema/payloads/$$dir; \
		openapi2jsonschema -block_additional_properties -go_constants -allow_null_values -spec ./swagger/payloads/$$dir/payload.yaml -out=./jsonschema/payloads/$$dir; \
  		echo generated $$dir; \
  	done

.PHONY: gen
gen:
	go generate ./...

.PHONY: build_tool
build_tool:
	go build -o ../bin/constanter ../cmd/constanter

.PHONY: build
build:
	go run .
