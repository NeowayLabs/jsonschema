package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func Check(data []byte, schema []byte) error {

	parsedData := map[string]interface{}{}
	parsedSchema := map[string]interface{}{}

	// TODO: check invalid data
	json.Unmarshal(data, &parsedData)
	// TODO: check invalid schema
	json.Unmarshal(schema, &parsedSchema)

	if len(parsedData) == 0 {
		return errors.New("input data is empty")
	}

	if len(parsedSchema) == 0 {
		return errors.New("input schema is empty")
	}

	for field, value := range parsedData {
		if parsedSchema[field] == nil {
			return errors.New("TODO:1")
		}

		s := parsedSchema[field]
		t := s.(map[string]interface{})["type"]
		if s == nil || t == nil {
			return errors.New("TODO:2")
		}

		if t.(string) == "object" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateObject(value.(map[string]interface{}), o.(map[string]interface{}))
			}
			return errors.New("TODO:3")
		} else if t.(string) == "array" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateArray(value.([]interface{}), o.(map[string]interface{}))
			}
			return errors.New("TODO:4")
		} else {
			valueType := reflect.TypeOf(value).String()
			expectedType := typeMapping(t.(string))

			if valueType != expectedType {
				return fmt.Errorf(
					"expected type[%s] got type[%s] value[%s]",
					expectedType,
					valueType,
					value,
				)
			}
		}

	}

	return nil
}

func validateArray(values []interface{}, schema map[string]interface{}) error {

	for _, data := range values {
		t := schema["type"]
		if t != nil {

			if t.(string) == "object" {
				o := schema["format"]
				if o != nil && reflect.TypeOf(data).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateObject(data.(map[string]interface{}), o.(map[string]interface{}))
				}
				return errors.New("TODO: 1")
			}

			if t.(string) == "array" {
				o := schema["format"]
				if o != nil && reflect.TypeOf(data).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateArray(data.([]interface{}), o.(map[string]interface{}))
				}

				return errors.New("TODO: 2")
			}

			return errors.New("TODO: 3")
		}
	}

	// TODO: test
	return nil
}

func validateObject(data, schema map[string]interface{}) error {

	for field, value := range data {

		if schema[field] == nil {
			return errors.New("TODO:1")
		}

		s := schema[field]
		t := s.(map[string]interface{})["type"]
		if s == nil || t == nil {
			return errors.New("TODO:2")
		}

		if t.(string) == "object" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateObject(value.(map[string]interface{}), o.(map[string]interface{}))
			}
			return errors.New("TODO:3")
		}

		if t.(string) == "array" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateArray(value.([]interface{}), o.(map[string]interface{}))
			}
			return errors.New("TODO:4")
		}

		if reflect.TypeOf(value).String() != typeMapping(t.(string)) {
			return errors.New("TODO:5")
		}
	}

	return nil
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
