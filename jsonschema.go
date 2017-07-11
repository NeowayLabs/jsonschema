package jsonschema

import (
	"encoding/json"
	"errors"
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
			// TODO: test
			return nil
		}

		s := parsedSchema[field]
		t := s.(map[string]interface{})["type"]
		if s == nil || t == nil {
			// TODO: test
			return nil
		}

		if t.(string) == "object" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateObject(value.(map[string]interface{}), o.(map[string]interface{}))
			}
			// TODO: test
			return nil
		} else if t.(string) == "array" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateArray(value.([]interface{}), o.(map[string]interface{}))
			}
			// TODO: Test
			return nil
		} else {
			if reflect.TypeOf(value).String() != typeMapping(t.(string)) {
				return errors.New("TODO: Improve error message")
			}
		}

	}

	// FIXME: Empty data + schema = ok. This seems wrong
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
				} else {
					// TODO: test
					return nil
				}
			} else if t.(string) == "array" {
				o := schema["format"]
				if o != nil && reflect.TypeOf(data).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
					return validateArray(data.([]interface{}), o.(map[string]interface{}))
				} else {
					// TODO: test
					return nil
				}
			} else {
				// TODO: test
				return nil
			}
		}
	}

	// TODO: test
	return nil
}

func validateObject(data, schema map[string]interface{}) error {

	for field, value := range data {

		if schema[field] == nil {
			// TODO: Test
			return nil
		}

		s := schema[field]
		t := s.(map[string]interface{})["type"]
		if s == nil || t == nil {
			// TODO: test
			return nil
		}

		if t.(string) == "object" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "map[string]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateObject(value.(map[string]interface{}), o.(map[string]interface{}))
			}
			// TODO: test
			return nil
		}

		if t.(string) == "array" {
			o := s.(map[string]interface{})["format"]
			if o != nil && reflect.TypeOf(value).String() == "[]interface {}" && reflect.TypeOf(o).String() == "map[string]interface {}" {
				return validateArray(value.([]interface{}), o.(map[string]interface{}))
			}
			// TODO: test
			return nil
		}

		if reflect.TypeOf(value).String() != typeMapping(t.(string)) {
			return errors.New("TODO:ERROR MESSAGE")
		}
	}

	// TODO: test
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
