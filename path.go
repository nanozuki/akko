package ononoki

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nanozuki/ononoki/typ"
)

type PathBuilder struct {
	path string
	item *openapi3.PathItem
}

func CONNECT(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func DELETE(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func GET(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func HEAD(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func OPTIONS(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func PATCH(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func POST(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func PUT(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func TRACE(id string, path string) *PathBuilder {
	return &PathBuilder{
		path: path,
		item: &openapi3.PathItem{},
	}
}

func (b *PathBuilder) Summary(summary string) *PathBuilder {
	b.item.Summary = summary
	return b
}

func (b *PathBuilder) Description(description string) *PathBuilder {
	b.item.Description = description
	return b
}

func (b *PathBuilder) Servers(servers ...*ServerBuilder) *PathBuilder {
	for _, server := range servers {
		b.item.Servers = append(b.item.Servers, server.server)
	}
	return b
}

func (b *PathBuilder) Request(props ...typ.ParameterPropBuilder) *PathBuilder {
	return b
}

func (b *PathBuilder) Response(props ...typ.ModelPropBuilder) *PathBuilder {
	return b
}
