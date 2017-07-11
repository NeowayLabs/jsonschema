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

func checkObject(data map[string]interface{}, schema map[string]interface{}) error {
	//for field, value := range data {
	//}
	return nil
}
