PAYLOADS = $(shell cd ./swagger/payloads && find | grep -oP ^.*schema$)

.PHONY: all
all: gen_schemas gen build

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
.PHONY: build
build:
	go run .
