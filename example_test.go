package jsonschema_test

import (
	"fmt"

	"github.com/NeowayLabs/jsonschema"
)

func ExampleCheck() {
	data := []byte(`{
		"objectField": {
			"stringField" : "name"
		}
	}`)
	schema := []byte(`{
		"objectField": {
			"type" : "object",
			"format" : {
				"stringField" : {
					"type" : "string"
				}
			}
		}
	}`)

	err := jsonschema.Check(data, schema)
	if err == nil {
		fmt.Println("success")
	}
	// Output: success
}
