package akko

import "github.com/getkin/kin-openapi/openapi3"

type ServerBuilder struct {
	server *openapi3.Server
}

func Server(url string) *ServerBuilder {
	return &ServerBuilder{&openapi3.Server{URL: url}}
}

func (b *ServerBuilder) Description(desc string) *ServerBuilder {
	b.server.Description = desc
	return b
}

func (b *ServerBuilder) Variable(name string, v *ServerVariableBuilder) *ServerBuilder {
	b.server.Variables[name] = v.v
	return b
}

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
