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
	info.AttachToAPI(b)
	return b
}

func (b *OpenAPIBuilder) Server(servers ...*ServerBuilder) *OpenAPIBuilder {
	for _, s := range servers {
		s.AttachToAPI(b)
	}
	return b
}

func (b *OpenAPIBuilder) Path(path *PathBuilder) *OpenAPIBuilder {
	panic("not implemented")
}

func (b *OpenAPIBuilder) Security(provider string, scopes ...string) *OpenAPIBuilder {
	sr := openapi3.NewSecurityRequirement().Authenticate(provider, scopes...)
	b.api.Security.With(sr)
	return b
}

func (b *OpenAPIBuilder) Tag(tags ...*TagBuilder) *OpenAPIBuilder {
	for _, tag := range tags {
		tag.AttachToAPI(b)
	}
	return b
}
