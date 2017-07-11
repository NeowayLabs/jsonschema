// Package jsonschema defines all functions required to
// validate json data according to a schema language
// that is also represented using JSON.
//
// For the full spec of the schema language check the
// project page: https://github.com/NeowayLabs/jsonschema
package jsonschema

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

// Check will check the given data according to
// the provided schema. If the data matches the given schema it
// will return nil, otherwise an error with details on why
// the given data does not conform to the provided schema.
func Check(data []byte, schema []byte) error {

	// TODO: accumulate all errors on data instead of returning
	// the first found error (avoid ping/pong of errors).

	// TODO: test obligatory fields

	parsedData := map[string]interface{}{}
	parsedSchema := map[string]interface{}{}

	err := json.Unmarshal(data, &parsedData)
	if err != nil {
		return fmt.Errorf("error[%s] parsing non JSON data[%s]", err, string(data))
	}

	err = json.Unmarshal(schema, &parsedSchema)
	if err != nil {
		return fmt.Errorf("error[%s] parsing non JSON schema[%s]", err, string(schema))
	}

	if len(parsedData) == 0 {
		return errors.New("input data is empty")
	}

	if len(parsedSchema) == 0 {
		return errors.New("input schema is empty")
	}

	return checkObject(parsedData, parsedSchema)
}

type typeDescriptor struct {
	Type   string
	Format interface{}
}

type typechecker func(rawdata interface{}, rawformat interface{}) error

func checkObject(rawdata interface{}, rawformat interface{}) error {

	data, ok := rawdata.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected data to be an 'object', it is: %q", reflect.TypeOf(rawdata))
	}
	// TODO: handle rawformat is not object
	format := rawformat.(map[string]interface{})

	for field, value := range data {
		desc, err := parseFieldTypeDescriptor(format, field)
		if err != nil {
			return fmt.Errorf("error getting type descriptor for field[%s]: %s", field, err)
		}
		// TODO: handle unknown type
		checker, err := getchecker(desc.Type)
		if err != nil {
			return fmt.Errorf("error getting type checker for field[%s]: %s", field, err)
		}
		if err := checker(value, desc.Format); err != nil {
			return fmt.Errorf("error validating field[%s] value[%s]: %s", field, value, err)
		}
	}

	return nil
}

func checkString(rawdata interface{}, format interface{}) error {
	// TODO: implement support to format on strings

	_, ok := rawdata.(string)
	if !ok {
		return fmt.Errorf("expected string, got [%s]", reflect.TypeOf(rawdata))
	}
	return nil
}

func checkFloat(rawdata interface{}, format interface{}) error {
	_, ok := rawdata.(float64)
	if !ok {
		return fmt.Errorf("expected float, got [%s]", reflect.TypeOf(rawdata))
	}
	return nil
}

func checkInt(rawdata interface{}, format interface{}) error {
	// TODO
	return nil
}

func checkArray(rawdata interface{}, rawformat interface{}) error {
	data, ok := rawdata.([]interface{})
	if !ok {
		return fmt.Errorf("expected data to be an 'array', it is: %q", reflect.TypeOf(rawdata))
	}
	// TODO: handle rawformat is not object
	format := rawformat.(map[string]interface{})

	for _, value := range data {
		desc, err := parseTypeDescriptor(format)
		if err != nil {
			return fmt.Errorf("error parsing type descriptor from format[%s]: %s", format, err)
		}
		checker, err := getchecker(desc.Type)
		if err != nil {
			return fmt.Errorf("error getting type checker for type[%s]: %s", desc.Type, err)
		}
		if err := checker(value, desc.Format); err != nil {
			return fmt.Errorf("error validating value[%s]: %s", value, err)
		}
	}

	return nil
}

func getchecker(typename string) (typechecker, error) {
	switch typename {
	case "string":
		{
			return checkString, nil
		}
	case "float":
		{
			return checkFloat, nil
		}
	case "int":
		{
			return checkInt, nil
		}
	case "object":
		{
			return checkObject, nil
		}
	case "array":
		{
			return checkArray, nil
		}
	}

	return nil, fmt.Errorf("unknown type[%s]", typename)
}

func parseFieldTypeDescriptor(schema map[string]interface{}, field string) (typeDescriptor, error) {
	// TODO: handle field not found
	rawDescriptor, ok := schema[field]
	if !ok {
		return typeDescriptor{}, fmt.Errorf("unable to find [%s] in schema[%s]", field, schema)
	}

	return parseTypeDescriptor(rawDescriptor)
}

func parseTypeDescriptor(rawDescriptor interface{}) (typeDescriptor, error) {
	// TODO: handle descriptor of wrong type
	parsedDescriptor := rawDescriptor.(map[string]interface{})

	// TODO: handle type of wrong type =P
	rawType := parsedDescriptor["type"]
	parsedType := rawType.(string)

	// TODO: handle format of wrong type
	var parsedFormat map[string]interface{}
	rawFormat, ok := parsedDescriptor["format"]
	if ok {
		parsedFormat = rawFormat.(map[string]interface{})
	}

	return typeDescriptor{
		Type:   parsedType,
		Format: parsedFormat,
	}, nil
}
