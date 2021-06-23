package ononoki

import "github.com/getkin/kin-openapi/openapi3"

type TagBuilder struct {
	tag *openapi3.Tag
}

func Tag(name string) *TagBuilder {
	return &TagBuilder{&openapi3.Tag{Name: name}}
}

func (b *TagBuilder) Description(desc string) *TagBuilder {
	b.tag.Description = desc
	return b
}
