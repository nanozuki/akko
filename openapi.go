package ononoki

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

func (b *OpenAPIBuilder) NewServer(baseURL string) *ServerBuilder {
	server := &openapi3.Server{
		URL: baseURL,
	}
	b.api.Servers = append(b.api.Servers, server)
	return &ServerBuilder{
		server: server,
	}
}

func (b *OpenAPIBuilder) NewPath(path string) *PathBuilder {
	return &PathBuilder{
		opanapi: b.api,
		baseURL: path,
		items:   map[Method]*OperationBuilder{},
	}
}

func (b *OpenAPIBuilder) WithSecure(provider string, scopes ...string) *OpenAPIBuilder {
	sr := openapi3.NewSecurityRequirement().Authenticate(provider, scopes...)
	b.api.Security.With(sr)
	return b
}

func (b *OpenAPIBuilder) NewTag(name string) *TagBuilder {
	tag := &openapi3.Tag{
		Name: name,
	}
	b.api.Tags = append(b.api.Tags, tag)
	return &TagBuilder{
		tag: tag,
	}
}
