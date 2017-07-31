# Json Schema

[![GoDoc](https://godoc.org/github.com/NeowayLabs/jsonschema?status.svg)](https://godoc.org/github.com/NeowayLabs/jsonschema) 
[![Build Status](https://travis-ci.org/NeowayLabs/jsonschema.svg?branch=master)](https://travis-ci.org/NeowayLabs/jsonschema)

## Why

Well, why a schema language ? We have internal APIs that integrates with
multiple teams, and some APIs that are consumed by third parties. On
these scenarios working just with documents + examples was not being
enough, specially to avoid automated validation of JSON documents.

Since we already work ubiquitously with JSON and there is tools
to work with it in practically any language we also use JSON
to express the desired schema for JSON documents.

This has the same objective as [JSON-Schema](http://json-schema.org/),
we rolled out our own because we hope to keep it MUCH more simple
than that, specially because our scenario is more restricted, we do
not have the goal to provide a universal JSON schema language that
contemplates every possible JSON document in the world, we
are just focusing on our needs.

## Specification

The full specification of the schema language
can be found [here](./docs/spec.md).
