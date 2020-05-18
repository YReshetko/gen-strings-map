
## Use case example

This example is created based on openapi2jsonschema tool which generates go files with constants uof jsonschema to validate a payload of request\response.

### Run
###### NOTE 
>Before running the following command you need to make sure that `openapi2jsonschema` tool is installed on your environment

```shell script
make all
```

The command will generate jsonschema with schemas and generated file of constants. The constants will be verified in main function of the test program