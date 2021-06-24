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
	v.AttachToServer(b, name)
	return b
}

func (b *ServerBuilder) AttachToAPI(api *OpenAPIBuilder) *ServerBuilder {
	api.api.Servers = append(api.api.Servers, b.server)
	return b
}

func (b *ServerBuilder) AttachToOp(op *OperationBuilder) *ServerBuilder {
	servers := *op.operation.Servers
	servers = append(servers, b.server)
	op.operation.Servers = &servers
	return b
}
