package jsonschema

import (
	"fmt"
	"reflect"
)

func Check(data, schema map[string]interface{}) bool {

	for field, value := range data {
		fmt.Println("field:", field, ", Value:", value)
		fmt.Println("schema:", schema[field])

		s := schema[field]
		t := s.(map[string]interface{})["type"]
		if s != nil && t != nil {
			fmt.Println("value type", reflect.TypeOf(value).String())
			fmt.Println("schema type", t.(string))

			if t.(string) == "object" {
				o := s.(map[string]interface{})["format"]
				if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					if !validateObject(value.(map[string]interface{}), o.(map[string]interface{})) {
						return false
					}
				} else {
					return false
				}
			} else if t.(string) == "array" {
				o := s.(map[string]interface{})["format"]
				if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					if !validateArray(value.([]interface{}), o.(map[string]interface{})) {
						return false
					}
				} else {
					return false
				}
			} else {
				if reflect.TypeOf(value).String() != typeMapping(t.(string)) {
					return false
				}
			}
		} else {
			return false
		}

	}

	return true
}

func validateArray(values []interface{}, schema map[string]interface{}) bool {

	fmt.Println("data", values)
	fmt.Println("schema", schema)

	for _, data := range values {
		t := schema["type"]
		if t != nil {
			fmt.Println("Schema type", t.(string))

			if t.(string) == "object" {
				o := schema["format"]
				if o != nil && reflect.TypeOf(data).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateObject(data.(map[string]interface{}), o.(map[string]interface{}))
				} else {
					return false
				}
			} else if t.(string) == "array" {
				o := schema["format"]
				if o != nil && reflect.TypeOf(data).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateArray(data.([]interface{}), o.(map[string]interface{}))
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}

	return true
}

func validateObject(data, schema map[string]interface{}) bool {

	for field, value := range data {

		fmt.Println("field:", field, ", Value:", value)
		fmt.Println("schema:", schema[field])

		s := schema[field]
		t := s.(map[string]interface{})["type"]
		if s != nil && t != nil {
			fmt.Println("value type", reflect.TypeOf(value).String())
			fmt.Println("schema type", t.(string))

			if t.(string) == "object" {
				o := s.(map[string]interface{})["format"]
				if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateObject(value.(map[string]interface{}), o.(map[string]interface{}))
				} else {
					return false
				}
			} else if t.(string) == "array" {
				o := s.(map[string]interface{})["format"]
				if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateArray(value.([]interface{}), o.(map[string]interface{}))
				} else {
					return false
				}
			} else {
				if reflect.TypeOf(value).String() != typeMapping(t.(string)) {
					return false
				}
			}
		} else {
			return false
		}
	}

	return true
}

func typeMapping(t string) string {
	types := map[string]string{
		"float": "float64",
	}

	newType, found := types[t]
	if found {
		return newType
	}

	return t
}
