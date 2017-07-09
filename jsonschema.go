package jsonschema

import (
	"fmt"
	"reflect"
)

func Check(data, schema map[string]interface{}) bool {

	for field, value := range data {
		fmt.Println("Field:", field, ", Value:", value)
		fmt.Println("Schema:", schema[field])

		s := schema[field]
		if s != nil && s.(map[string]interface{})["type"] != nil {
			fmt.Println("Value type", reflect.TypeOf(value).String())
			fmt.Println("Schema type", s.(map[string]interface{})["type"].(string))

			if reflect.TypeOf(value).String() != s.(map[string]interface{})["type"].(string) {
				return false
			}
		}

	}

	return true
}
