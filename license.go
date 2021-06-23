package ononoki

import "github.com/getkin/kin-openapi/openapi3"

type LicenseBuilder struct {
	license *openapi3.License
}

func License(name string) *LicenseBuilder {
	return &LicenseBuilder{&openapi3.License{Name: name}}
}

func (b *LicenseBuilder) URL(url string) *LicenseBuilder {
	b.license.URL = url
	return b
}
