package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + pwd + "/schema/schema.json")
	documentLoader := gojsonschema.NewReferenceLoader("file:///" + pwd + "/schema/post.json")

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}
