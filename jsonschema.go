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
	// handle rawdata is not object
	data := rawdata.(map[string]interface{})
	// handle rawformat is not object
	format := rawformat.(map[string]interface{})

	for field, _ := range data {
		// handle error
		desc, _ := parseTypeDescriptor(format, field)
		// handle unknown type
		getchecker(desc.Type)

	}
	return nil
}

func checkString(rawdata interface{}, format interface{}) error {
	return nil
}

func getchecker(typename string) typechecker {
	switch typename {
	case "string":
		{
			return checkString
		}
	}
	return nil
}

func parseTypeDescriptor(schema map[string]interface{}, field string) (typeDescriptor, error) {
	// TODO: handle field not found
	rawDescriptor, ok := schema[field]
	if !ok {
		return typeDescriptor{}, fmt.Errorf("unable to find [%s] in schema[%s]", field, schema)
	}
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
