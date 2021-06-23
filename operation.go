package ononoki

import "github.com/getkin/kin-openapi/openapi3"

type OperationBuilder struct {
	operation *openapi3.Operation
}

func Operation(id string) *OperationBuilder {
	return &OperationBuilder{
		&openapi3.Operation{OperationID: id},
	}
}
