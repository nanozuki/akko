package desc

// Prop descript a property in open-api,
// it will be a field in request or response
type Prop struct {
	Name        string
	Type        *TypeInfo
	StructField string
	Attr        SwaggerAttr

	modelProp
	paramterProp
}

type ModelProp interface {
	IsModelProp()
}

type modelProp struct{}

func (p modelProp) IsModelProp() {}

type ParamterProp interface {
	IsParameterProp()
}

type paramterProp struct{}

func (p paramterProp) IsParameterProp() {}

type InvalidProp struct{}

func (p *Prop) In(loc Location) *ParamProp {
	return &ParamProp{
		Prop: p,
		In:   loc,
	}
}

// SwaggerAttr present properties schema in swagger.
type SwaggerAttr struct {
	Title       string
	Type        string
	Required    string
	Description string
}

/* Swagger properties (https://swagger.io/specification/#SchemaObject):

1. The following properties are taken directly from the JSON Schema definition
and follow the same specifications:
	[x] title
	[ ] multipleOf
	[ ] maximum
	[ ] exclusiveMaximum
	[ ] minimum
	[ ] exclusiveMinimum
	[ ] maxLength
	[ ] minLength
	[ ] pattern
	[ ] maxItems
	[ ] minItems
	[ ] uniqueItems
	[ ] maxProperties
	[ ] minProperties
	[x] required
	[ ] enum

2. The following properties are taken from the JSON Schema definition
but their definitions were adjusted to the OpenAPI Specification.
	[x] type
	[ ] allOf
	[ ] oneOf
	[ ] anyOf
	[ ] not
	[ ] items
	[x] properties
	[ ] additionalProperties
	[x] description
	[ ] format
	[ ] default
*/
