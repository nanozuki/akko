package akko

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nanozuki/akko/typ"
)

type OperationBuilder struct {
	operation *openapi3.Operation
}

func Op(id string) *OperationBuilder {
	return &OperationBuilder{
		&openapi3.Operation{OperationID: id},
	}
}

func (b *OperationBuilder) Request(props ...typ.ParameterPropBuilder) *OperationBuilder {
	return b
}

func (b *OperationBuilder) Response(props ...typ.ModelPropBuilder) *OperationBuilder {
	return b
}
