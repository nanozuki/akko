package akko

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type OpenAPIBuilder struct {
	api *openapi3.T
}

func OpenAPI(title string, version string) *OpenAPIBuilder {
	return &OpenAPIBuilder{
		api: &openapi3.T{
			OpenAPI: "v3.0.3",
			Info: &openapi3.Info{
				Title:   title,
				Version: version,
			},
		},
	}
}

func (b *OpenAPIBuilder) Info(info *InfoBuilder) *OpenAPIBuilder {
	info.info.Title = b.api.Info.Title
	info.info.Version = b.api.Info.Version
	b.api.Info = info.info
	return b
}

func (b *OpenAPIBuilder) Server(servers ...*ServerBuilder) *OpenAPIBuilder {
	for _, s := range servers {
		b.api.Servers = append(b.api.Servers, s.server)
	}
	return b
}

func (b *OpenAPIBuilder) Path(path *PathBuilder) *OpenAPIBuilder {
	b.api.Paths[path.path] = path.item
	return b
}

func (b *OpenAPIBuilder) Secure(provider string, scopes ...string) *OpenAPIBuilder {
	sr := openapi3.NewSecurityRequirement().Authenticate(provider, scopes...)
	b.api.Security.With(sr)
	return b
}

func (b *OpenAPIBuilder) Tag(tags ...*TagBuilder) *OpenAPIBuilder {
	for _, tag := range tags {
		b.api.Tags = append(b.api.Tags, tag.tag)
	}
	return b
}

func (b *OpenAPIBuilder) Run() {
	panic("not implemented")
}
