<!-- mdtocstart -->

# Table of Contents

- [JSON Schema Specification](#json-schema-specification)
    - [String](#string)
    - [Float](#float)
    - [Int](#int)
    - [Object](#object)
    - [Array](#array)

<!-- mdtocend -->

# JSON Schema Specification

Here is defined a schema language that uses JSON to
define the structure of JSON objects.

The first concept you need to understand is the type
descriptor. Since JSON objects are composed by a map
between field names and values we start with the idea
of describing the type of individual fields.

A type descriptor is a JSON object like this:

```json
{
    "type" : "string",
    "required": true,
    "format" : "YYYY-MM-DD"
}
```

This descriptor is defining a field with type string
that is required and has the format **YYYY-MM-DD**.

The full schema for the field **name** using the
descriptor above:

```json
{
    "name": {
        "type" : "string",
        "required": true,
        "format" : "YYYY-MM-DD"
    }
}
```

The **type** and **required** fields are mandatory.
The **format** field may be omitted depending on the type.

Now that you understand the basic concept of a
field descriptor lets go through a complete list
of all available type descriptors.

## String

The string descriptor is a JSON object like this:

```json
{
    "type" : "string",
    "required": true,
    "format" : "YYYY-MM-DD"
}
```

Where **format** is optional. Here is a full example of multiple
string fields:

```json
{
    "name": {
        "type" : "string",
        "required": true,
    },
    "lastname": {
        "type" : "string",
        "required": true,
    }
    "birthday": {
        "type" : "string",
        "required": false,
        "format" : "YYYY-MM-DD"
    }
}
```

## Float

The float descriptor represents float numbers,
it is a JSON object like this:

```json
{
    "type" : "float",
    "required": true
}
```

There is no **format** option for float.
Here is a full example of multiple float fields:

```json
{
    "salary": {
        "type" : "float",
        "required": true,
    },
    "debt": {
        "type" : "float",
        "required": false,
    },
}
```

## Int

The int descriptor represents int numbers,
it is a JSON object like this:

```json
{
    "type" : "int",
    "required": true
}
```

There is no **format** option for int.
Here is a full example of multiple int fields:

```json
{
    "age": {
        "type" : "int",
        "required": true,
    },
    "count": {
        "type" : "int",
        "required": false,
    }
}
```

## Object

The object descriptor represents nesting on a JSON object, it will
happen when a field maps to another object.

This situation is naturally a recursion, the recursion will be
expressed on the **format** field, which in this case is itself
a JSON schema describing the nested object:

```json
{
    "type" : "number",
    "required": true,
    "format" : { ... }
}
```

The **format** option is required in this case since it will
define the schema of the nested object.

Here is a full example of a person with a nested object
containing the name and age:

```json
{
    "person" : {
        "type": "object",
        "required": true,
        "format": {
            "name": {
                "type" : "string",
                "required": true,
            },
            "age": {
                "type" : "int",
                "required": true,
            }
        }
    }
}
```

The JSON bellow would conform with the schema
defined above:

```json
{
    "person" : {
        "name": "leo",
        "age": 45
    }
}
```

## Array

An array can be a collection of single valued types
like strings, floats and ints or from another collection
types as objects or even arrays (a recursion).

To express this kind of complexity we define an array
using the **format** field to express recursion, just
as it is done with the object:

```json
{
    "type": "array",
    "required": true,
    "format": {
        "type" : "string"
    }
}
```

The array descriptor **format** field must be an object that
is the type descriptor of the values expected to be
inside the array.

It is VERY unusual to have an array with
more than one type inside (it is usually a terrible idea), so
we disallow this on the schema description, an array can have
one type of data inside it.

The type descriptor that is inside the array do not need the
**required** field since it represent values that will be
inside an array (they are not a field).

Some examples, first an array of integers:

```json
{
    "integersArray" : {
        "type": "array",
        "required": true,
        "format": { "type" : "int" }
    }
}
```

Array of strings:

```json
{
    "stringArray" : {
        "type": "array",
        "required": true,
        "format": { "type" : "string" }
    }
}
```

Array of objects:

```json
{
    "objectArray" : {
        "type": "array",
        "required": true,
        "format": {
            "type" : "object"
            "format": {
                "name" : {
                    "type": "string",
                    "required" : true
                },
                "age" : {
                    "type": "int",
                    "required" : true
                }
            }
        }
    }
}
```

Array of object array:

```json
{
    "objectArray" : {
        "type": "array",
        "required": true,
        "format": {
            "type": "array",
            "format" : {
                "type" : "object"
                "format": {
                    "name" : {
                        "type": "string",
                        "required" : true
                    },
                    "age" : {
                        "type": "int",
                        "required" : true
                    }
                }
            }
        }
    }
}
```
