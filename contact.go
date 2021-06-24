package akko

import "github.com/getkin/kin-openapi/openapi3"

type ContactBuilder struct {
	contact *openapi3.Contact
}

func Contact() *ContactBuilder {
	return &ContactBuilder{&openapi3.Contact{}}
}

func (b *ContactBuilder) Name(name string) *ContactBuilder {
	b.contact.Name = name
	return b
}

func (b *ContactBuilder) URL(url string) *ContactBuilder {
	b.contact.URL = url
	return b
}

func (b *ContactBuilder) Email(email string) *ContactBuilder {
	b.contact.Email = email
	return b
}
