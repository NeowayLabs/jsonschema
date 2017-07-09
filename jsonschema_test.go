package jsonschema

import (
	"testing"
)

func TestCheckUsingFieldStringShouldReturnTrue(t *testing.T) {

	schema := map[string]interface{}{
		"stringField": map[string]interface{}{
			"type": "string",
		},
	}

	data := map[string]interface{}{
		"stringField": "name",
	}

	expected := true
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldStringShouldReturnFalseWhenTypeIsntString(t *testing.T) {

	schema := map[string]interface{}{
		"stringField": map[string]interface{}{
			"type": "string",
		},
	}

	data := map[string]interface{}{
		"stringField": 1,
	}

	expected := false
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldIntShouldReturnTrue(t *testing.T) {

	schema := map[string]interface{}{
		"intField": map[string]interface{}{
			"type": "int",
		},
	}

	data := map[string]interface{}{
		"intField": 1,
	}

	expected := true
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldIntShouldReturnFalseWhenTypeIsntInt(t *testing.T) {

	schema := map[string]interface{}{
		"intField": map[string]interface{}{
			"type": "int",
		},
	}

	data := map[string]interface{}{
		"intField": "1",
	}

	expected := false
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldFloatShouldReturnTrue(t *testing.T) {

	schema := map[string]interface{}{
		"intField": map[string]interface{}{
			"type": "float",
		},
	}

	data := map[string]interface{}{
		"intField": 1.0,
	}

	expected := true
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldFloatShouldReturnFalseWhenTypeIsntFloat(t *testing.T) {

	schema := map[string]interface{}{
		"intField": map[string]interface{}{
			"type": "float",
		},
	}

	data := map[string]interface{}{
		"intField": "1",
	}

	expected := false
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingInvalidSchemaMustReturnFalse(t *testing.T) {

	schema := map[string]interface{}{
		"intField": map[string]interface{}{},
	}

	data := map[string]interface{}{
		"intField": "1",
	}

	expected := false
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
