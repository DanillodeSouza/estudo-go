package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + pwd + "/schema/schema.json")

	post := `{
		"title": "Desenvolvedor Gopher",
		"role": {
			"id" : 446
		},
		"cities": [{
			"id": 1,
			"quantity": 2
		},{
			"id": 3,
			"quantity": 10
		}],
		"salaryNegotiable": false
	}`

	stringLoader := gojsonschema.NewStringLoader(post)

	result, err := gojsonschema.Validate(schemaLoader, stringLoader)
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
