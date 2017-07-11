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
			success: true, //TODO
		},
		Scenario{
			name: "UnknowDataField",
			data: `{
				"unknow" : 1
			}`,
			schema: `{
				"intfield" : {
					"type": "int"
				}
			}`,
			success: true, //TODO
		},
		Scenario{
			name: "EmptyData",
			data: `{}`,
			schema: `{
				"floatField" : {
					"type": "float"
				}
			}`,
			success: false,
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
			name: "WrongString",
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
			name: "WrongArray",
			data: `{
				"arrayField" : 1
			}`,
			schema: `{
				"arrayField" : {
					"type": "array",
					"format" : {
						"type" : "string"
					}
				}
			}`,
			success: true, // TODO
		},
		Scenario{
			name: "WrongNestedArray",
			data: `{
				"arrayField" : [ [ { "stringField" : 1 } ] ]
			}`,
			schema: `{
				"arrayField" : {
					"type": "array",
					"format" : {
						"type" : "array",
						"format" : {
							"type" : "string"
						}
					}
				}
			}`,
			success: true, //TODO
		},
		Scenario{
			name: "WrongObjectInsideNestedArray",
			data: `{
				"arrayField" : [ [ { "stringField" : 1 } ] ]
			}`,
			schema: `{
				"arrayField" : {
					"type": "array",
					"format" : {
						"type" : "array",
						"format" : {
							"type" : "object",
							"format" : {
								"stringField" : {
									"type" : "string"
								}
							}
						}
					}
				}
			}`,
			success: true, //TODO
		},
		Scenario{
			name: "WrongFloat",
			data: `{
				"floatField" : "lala"
			}`,
			schema: `{
				"floatField" : {
					"type": "float"
				}
			}`,
			success: true, // TODO
		},
		Scenario{
			name: "WrongObject",
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
			success: true, //TODO
		},
		Scenario{
			name: "corruptedData",
			data: `{
				objectField"="wrong"
			`,
			schema: `{
				"stringField": {
					"type" : "string"
				}
			}`,
			success: false,
		},
		Scenario{
			name: "corruptedSchema",
			data: `{
				"objectField": {
					"nestedFloat" : 1.1
				}
			}`,
			schema: `
				"objectField": {
					type" : "object",
					}
			}`,
			success: false,
		},
		Scenario{
			name: "WrongNestedFloat",
			data: `{
				"objectField": {
					"nestedFloat" : "wrong"
				}
			}`,
			schema: `{
				"objectField": {
					"type" : "object",
					"format" : {
						"nestedFloat" : {
							"type" : "float"
						}
					}
				}
			}`,
			success: true, // TODO
		},
		Scenario{
			name: "WrongNestedObject",
			data: `{
				"objectField": {
					"nestedObject" : "wrong"
				}
			}`,
			schema: `{
				"objectField": {
					"type" : "object",
					"format" : {
						"nestedObject" : {
							"type" : "object",
							"format" : {
								"s" : {
									"type" : "string"
								}
							}
						}
					}
				}
			}`,
			success: true, // TODO
		},
	}

	for _, scenario := range scenarios {
		testScenario(t, scenario)
	}
}

func TestSuccessOn(t *testing.T) {

	scenarios := []Scenario{
		Scenario{
			name: "Int",
			data: `{
				"intField": 1
			}`,
			schema: `{
				"intField": {
					"type" : "int"
				}
			}`,
			success: true,
		},
		Scenario{
			name: "Object",
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
			name: "ArrayOfObjects",
			data: `{
				"arrayField": [ {"stringField":"hi"} ]
			}`,
			schema: `{
				"arrayField": {
					"type" : "array",
					"format" : {
						"type" : "object",
						"format" : {
							"stringField" : {
								"type" : "string"
							}
						}
					}
				}
			}`,
			success: true,
		},
		Scenario{
			name: "NestedArrayOfObjects",
			data: `{
				"arrayField": [ [ {"stringField":"hi"} ] ]
			}`,
			schema: `{
				"arrayField": {
					"type" : "array",
					"format" : {
						"type" : "array",
						"format" : {
							"type" : "object",
							"format" : {
								"stringField" : {
									"type" : "string"
								}
							}
						}
					}
				}
			}`,
			success: true,
		},
		Scenario{
			name: "NestedObject",
			data: `{
				"objectField": {
					"nestedObject" : {
						"stringField" : "name"
					}
				}
			}`,
			schema: `{
				"objectField": {
					"type" : "object",
					"format" : {
						"nestedObject" : {
							"type" : "object",
							"format" : {
								"stringField" : {
									"type" : "string"
								}
							}
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
