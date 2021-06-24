package akko

import "github.com/getkin/kin-openapi/openapi3"

type ServerVariableBuilder struct {
	v *openapi3.ServerVariable
}

func ServerVariable(defaultValue string) *ServerVariableBuilder {
	return &ServerVariableBuilder{&openapi3.ServerVariable{Default: defaultValue}}
}

func (b *ServerVariableBuilder) Enum(values ...string) *ServerVariableBuilder {
	b.v.Enum = values
	return b
}

func (b *ServerVariableBuilder) Description(desc string) *ServerVariableBuilder {
	b.v.Description = desc
	return b
}

func (b *ServerVariableBuilder) AttachToServer(server *ServerBuilder, name string) *ServerVariableBuilder {
	server.server.Variables[name] = b.v
	return b
}
