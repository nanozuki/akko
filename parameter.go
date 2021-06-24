package akko

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nanozuki/akko/desc"
)

type ParameterBuilder struct {
	p *openapi3.Parameter
}

func Parameter(name string, in desc.Location) *ParameterBuilder {
	b := &ParameterBuilder{
		p: &openapi3.Parameter{
			Name: name,
			In:   string(in),
		},
	}
	if in == desc.InPath {
		b.p.Required = true
	}
	return b
}
