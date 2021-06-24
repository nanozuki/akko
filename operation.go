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

func (b *OperationBuilder) Tags(tags ...string) *OperationBuilder {
	b.operation.Tags = append(b.operation.Tags, tags...)
	return b
}

func (b *OperationBuilder) TagObjects(tags ...*TagBuilder) *OperationBuilder {
	for _, tag := range tags {
		tag.AttachToOp(b)
	}
	return b
}

func (b *OperationBuilder) Summary(summary string) *OperationBuilder {
	b.operation.Summary = summary
	return b
}

func (b *OperationBuilder) Description(description string) *OperationBuilder {
	b.operation.Description = description
	return b
}

func (b *OperationBuilder) Parameter(parameters ...typ.ParameterPropBuilder) *OperationBuilder {
	panic("not implemented")
	// return b
}

func (b *OperationBuilder) Request(props ...typ.ModelPropBuilder) *OperationBuilder {
	panic("not implemented")
	// return b
}

func (b *OperationBuilder) Response(props ...typ.ModelPropBuilder) *OperationBuilder {
	panic("not implemented")
	// return b
}

func (b *OperationBuilder) Callbacks(callbacks ...CallbackBuilder) *OperationBuilder {
	for _, c := range callbacks {
		c.AttachToOp(b)
	}
	return b
}

func (b *OperationBuilder) Deprecated(deprecated bool) *OperationBuilder {
	b.operation.Deprecated = deprecated
	return b
}

func (b *OperationBuilder) Security(provider string, scopes ...string) *OperationBuilder {
	sr := openapi3.NewSecurityRequirement().Authenticate(provider, scopes...)
	b.operation.Security.With(sr)
	return b
}

func (b *OperationBuilder) Servers(servers ...*ServerBuilder) *OperationBuilder {
	for _, s := range servers {
		s.AttachToOp(b)
	}
	return b
}
