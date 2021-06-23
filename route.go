package ononoki

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/nanozuki/ononoki/typ"
)

func Resquest(props ...typ.ParameterPropBuilder) {}
func Response(props ...typ.ModelPropBuilder)     {}

type OpenAPIBuilder struct {
	api openapi3.T
}

func OpenAPI(info *InfoBuilder) *OpenAPIBuilder {
	return &OpenAPIBuilder{
		api: openapi3.T{
			OpenAPI: "v3.0.3",
		},
	}
}

type InfoBuilder struct {
	info openapi3.Info
}

func Info(title string, version string) *InfoBuilder {
	return &InfoBuilder{
		info: openapi3.Info{
			Title:   title,
			Version: version,
		},
	}
}
