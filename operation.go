package ononoki

import "github.com/getkin/kin-openapi/openapi3"

type OperationBuilder struct {
	operation *openapi3.Operation
	url       string
	method    Method
}
