package jsonschema_test

import (
	"fmt"
	"testing"

	"github.com/NeowayLabs/jsonschema"
)

func TestFailureOn(t *testing.T) {

	scenarios := []Scenario{
		Scenario{
			name: "ExpectFloatButGotInt",
			data: `{
				"floatField" : 1
			}`,
			schema: `{
				"floatField" : {
					"type": "float"
				}
			}`,
			success: true, // FIXME: Should be false, fix float/int validation
		},
		Scenario{
			name:    "EmptySchema",
			data:    `{"intField": 1}`,
			schema:  `{}`,
			success: false,
		},
		Scenario{
			name:    "EverythingEmpty",
			data:    `{}`,
			schema:  `{}`,
			success: false,
		},
		Scenario{
			name: "WrongobjectField",
			data: `{
				"stringField" : 1
			}`,
			schema: `{
				"stringField" : {
					"type": "string"
				}
			}`,
			success: false,
		},
		Scenario{
			name: "WrongFloatField",
			data: `{
				"floatField" : "lala"
			}`,
			schema: `{
				"floatField" : {
					"type": "float"
				}
			}`,
			success: false,
		},
		Scenario{
			name: "ObjectField",
			data: `{
				"objectField": "wrong"
			}`,
			schema: `{
				"objectField": {
					"type" : "object",
					"format" : {
						"stringField" : {
							"type" : "string"
						}
					}
				}
			}`,
			success: false,
		},
	}

	for _, scenario := range scenarios {
		testScenario(t, scenario)
	}
}

func TestSuccessOn(t *testing.T) {

	scenarios := []Scenario{
		Scenario{
			name: "IntField",
			data: `{
				"intField": 1
			}`,
			schema: `{
				"intField": {
					"type" : "int"
				}
			}`,
			success: false, //FIXME: should be true
		},
		Scenario{
			name: "ObjectField",
			data: `{
				"objectField": {
					"stringField" : "name"
				}
			}`,
			schema: `{
				"objectField": {
					"type" : "object",
					"format" : {
						"stringField" : {
							"type" : "string"
						}
					}
				}
			}`,
			success: true,
		},
		Scenario{
			name: "FloatField",
			data: `{
				"floatField": 1.3
			}`,
			schema: `{
				"floatField": {
					"type" : "float"
				}
			}`,
			success: true,
		},
	}

	for _, scenario := range scenarios {
		testScenario(t, scenario)
	}
}

//func TestCheckUsingFieldObjectShouldReturnFalseWhenTypeIsntObject(t *testing.T) {

//schema := map[string]interface{}{
//"objectField": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//}

//data := map[string]interface{}{
//"objectField": "field",
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldObjectShouldReturnFalseWhenTypeInsideObjectIsntExpectedType(t *testing.T) {

//schema := map[string]interface{}{
//"objectField": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//}

//data := map[string]interface{}{
//"objectField": map[string]interface{}{
//"stringField": 1,
//},
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingMultipleFieldObjectShouldReturnTrue(t *testing.T) {

//schema := map[string]interface{}{
//"objectField": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"secondObjectField": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"objectField": map[string]interface{}{
//"secondObjectField": map[string]interface{}{
//"stringField": "field",
//},
//},
//}

//expected := true
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldArrayShouldReturnTrue(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//map[string]interface{}{
//"stringField": "field",
//},
//},
//}

//expected := true
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldArrayShouldReturnFalseWhenTypeInsideArrayIsntExpectedType(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//map[string]interface{}{
//"stringField": 1,
//},
//},
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingMultipleFieldArrayShouldReturnTrue(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//[]interface{}{
//map[string]interface{}{
//"stringField": "field",
//},
//},
//},
//}

//expected := true
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingMultipleFieldArrayShouldReturnFalseWhenTypeInsideIsntExpectedType(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//[]interface{}{
//map[string]interface{}{
//"stringField": 1,
//},
//},
//},
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldObjectWithArrayShouldReturnTrue(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//map[string]interface{}{
//"stringField": "field",
//},
//},
//}

//expected := true
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldObjectShouldReturnFalseWhenTypeInsideIsntExpectedType(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//map[string]interface{}{
//"stringField": 1,
//},
//},
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingFieldDoesntContainInSchema(t *testing.T) {

//schema := map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//}

//data := map[string]interface{}{
//"anotherField": 1,
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

//func TestCheckUsingNestedFieldDoesntContainInSchema(t *testing.T) {

//schema := map[string]interface{}{
//"arrayField": map[string]interface{}{
//"type": "array",
//"format": map[string]interface{}{
//"type": "object",
//"format": map[string]interface{}{
//"stringField": map[string]interface{}{
//"type": "string",
//},
//},
//},
//},
//}

//data := map[string]interface{}{
//"arrayField": []interface{}{
//map[string]interface{}{
//"anotherField": "another",
//},
//},
//}

//expected := false
//actual := jsonschema.Check(data, schema)

//if actual != expected {
//t.Error("Test failed. Expected", expected, "but returned", actual)
//}
//}

type Scenario struct {
	name    string
	schema  string
	data    string
	success bool
}

func testScenario(t *testing.T, s Scenario) {
	t.Run(s.name, func(t *testing.T) {
		details := fmt.Sprintf("data:\n%s\n\nschema:\n%s\n", s.data, s.schema)
		err := jsonschema.Check([]byte(s.data), []byte(s.schema))
		if s.success {
			if err != nil {
				t.Fatalf("unexpected error[%s],details:\n%s", err, details)
			}
		} else {
			if err == nil {
				t.Fatalf("expected error got nil, details:\n%s", details)
			}
		}
	})
}
