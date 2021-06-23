package ononoki

import (
	ospath "path"

	"github.com/getkin/kin-openapi/openapi3"
)

type PathBuilder struct {
	opanapi *openapi3.T
	baseURL string
	items   map[Method]*OperationBuilder
}

func (b *PathBuilder) Path(path string) *PathBuilder {
	return &PathBuilder{
		opanapi: b.opanapi,
		baseURL: ospath.Join(b.baseURL, path),
		items:   map[Method]*OperationBuilder{},
	}
}
