package akko

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

func (b *TagBuilder) AttachToAPI(api *OpenAPIBuilder) *TagBuilder {
	api.api.Tags = append(api.api.Tags, b.tag)
	return b
}

func (b *TagBuilder) AttachToOp(op *OperationBuilder) *TagBuilder {
	op.operation.Tags = append(op.operation.Tags, b.tag.Name)
	return b
}
