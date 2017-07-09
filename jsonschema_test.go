package jsonschema

import (
	"testing"
)

func TestCheckUsingFieldStringShouldReturnTrueWhenTypeIsString(t *testing.T) {

	schema := map[string]interface{}{
		"name": map[string]interface{}{
			"type": "string",
		},
	}

	data := map[string]interface{}{
		"name": "name",
	}

	expected := true
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}

func TestCheckUsingFieldStringShouldReturnFalseWhenTypeIsntString(t *testing.T) {

	schema := map[string]interface{}{
		"name": map[string]interface{}{
			"type": "string",
		},
	}

	data := map[string]interface{}{
		"name": 1,
	}

	expected := false
	actual := Check(data, schema)

	if actual != expected {
		t.Error("Test failed. Expected", expected, "but returned", actual)
	}
}
