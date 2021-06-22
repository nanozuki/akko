package prop

import "github.com/nanozuki/ononoki/desc"

type BoolBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *BoolBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *BoolBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *BoolBuilder) StructTag(key, value string) *BoolBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *BoolBuilder) Validator(v ...string) *BoolBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *BoolBuilder) Title(title string) *BoolBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *BoolBuilder) Description(desc string) *BoolBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *BoolBuilder) Example(example string) *BoolBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *BoolBuilder) Required() *BoolBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to Bool.
func (b *BoolBuilder) GoType(i interface{}) *BoolBuilder {
	b.goType(i, boolType)
	return b
}

// InPath to declare the localtion of this field is query string.
func (b *BoolBuilder) InQuery() *PBoolBuilder {
	b.inQuery()
	return (*PBoolBuilder)(b)
}

type StringBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *StringBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *StringBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *StringBuilder) StructTag(key, value string) *StringBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *StringBuilder) Validator(v ...string) *StringBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *StringBuilder) Title(title string) *StringBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *StringBuilder) Description(desc string) *StringBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *StringBuilder) Example(example string) *StringBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *StringBuilder) Required() *StringBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to String.
func (b *StringBuilder) GoType(i interface{}) *StringBuilder {
	b.goType(i, stringType)
	return b
}

// InPath to declare the localtion of this field is url path.
func (b *StringBuilder) InPath() *PStringBuilder {
	b.inPath()
	return (*PStringBuilder)(b)
}

// InPath to declare the localtion of this field is query string.
func (b *StringBuilder) InQuery() *PStringBuilder {
	b.inQuery()
	return (*PStringBuilder)(b)
}

// MaximumLength set maximum length of string.
func (b *StringBuilder) MaximumLength(max uint) *StringBuilder {
	b.maximumLength(max)
	return b
}

// MinimumLength set maximum length of string.
func (b *StringBuilder) MinimumLength(min uint) *StringBuilder {
	b.minimumLength(min)
	return b
}

// Pattern to set value pattern with regular expression.
// As go-validator not support regular expression will (because some the performance reason),
// You should set a go-validator struct filed.
func (b *StringBuilder) Pattern(regex string, validators ...string) *StringBuilder {
	b.pattern(regex, validators...)
	return b
}

// Enum to set the valid values for string property.
func (b *StringBuilder) Enum(values ...string) *StringBuilder {
	b.enum(values...)
	return b
}

type IntBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *IntBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *IntBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *IntBuilder) StructTag(key, value string) *IntBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *IntBuilder) Validator(v ...string) *IntBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *IntBuilder) Title(title string) *IntBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *IntBuilder) Description(desc string) *IntBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *IntBuilder) Example(example string) *IntBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *IntBuilder) Required() *IntBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to Int.
func (b *IntBuilder) GoType(i interface{}) *IntBuilder {
	b.goType(i, intType)
	return b
}

// InPath to declare the localtion of this field is url path.
func (b *IntBuilder) InPath() *PIntBuilder {
	b.inPath()
	return (*PIntBuilder)(b)
}

// InPath to declare the localtion of this field is query string.
func (b *IntBuilder) InQuery() *PIntBuilder {
	b.inQuery()
	return (*PIntBuilder)(b)
}

// Maximum set the maximum value of this property.
func (b *IntBuilder) Maximum(max int) *IntBuilder {
	b.maximum(&max, nil)
	return b
}

// Minimum set the minimum value of this property.
func (b *IntBuilder) Minimum(min int) *IntBuilder {
	b.minimum(&min, nil)
	return b
}

type FloatBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *FloatBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *FloatBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *FloatBuilder) StructTag(key, value string) *FloatBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *FloatBuilder) Validator(v ...string) *FloatBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *FloatBuilder) Title(title string) *FloatBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *FloatBuilder) Description(desc string) *FloatBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *FloatBuilder) Example(example string) *FloatBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *FloatBuilder) Required() *FloatBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to Float.
func (b *FloatBuilder) GoType(i interface{}) *FloatBuilder {
	b.goType(i, floatType)
	return b
}

// InPath to declare the localtion of this field is url path.
func (b *FloatBuilder) InPath() *PFloatBuilder {
	b.inPath()
	return (*PFloatBuilder)(b)
}

// InPath to declare the localtion of this field is query string.
func (b *FloatBuilder) InQuery() *PFloatBuilder {
	b.inQuery()
	return (*PFloatBuilder)(b)
}

// Maximum set the maximum value of this property.
func (b *FloatBuilder) Maximum(max float64) *FloatBuilder {
	b.maximum(nil, &max)
	return b
}

// Minimum set the minimum value of this property.
func (b *FloatBuilder) Minimum(min float64) *FloatBuilder {
	b.minimum(nil, &min)
	return b
}

type ArrayBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *ArrayBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *ArrayBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *ArrayBuilder) StructTag(key, value string) *ArrayBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *ArrayBuilder) Validator(v ...string) *ArrayBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *ArrayBuilder) Title(title string) *ArrayBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *ArrayBuilder) Description(desc string) *ArrayBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *ArrayBuilder) Example(example string) *ArrayBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *ArrayBuilder) Required() *ArrayBuilder {
	b.required()
	return b
}

// MaximumItems to set the maximum count of array items.
func (b *ArrayBuilder) MaximumItems(max uint) *ArrayBuilder {
	b.maximumItems(max)
	return b
}

// MinimumItems to set the minimum count of array items.
func (b *ArrayBuilder) MinimumItems(min uint) *ArrayBuilder {
	b.minimumItems(min)
	return b
}

// Unique to mark items in array is unique.
func (b *ArrayBuilder) Unique() *ArrayBuilder {
	b.unique()
	return b
}

type ObjectBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *ObjectBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *ObjectBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *ObjectBuilder) StructTag(key, value string) *ObjectBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *ObjectBuilder) Validator(v ...string) *ObjectBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *ObjectBuilder) Title(title string) *ObjectBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *ObjectBuilder) Description(desc string) *ObjectBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *ObjectBuilder) Example(example string) *ObjectBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *ObjectBuilder) Required() *ObjectBuilder {
	b.required()
	return b
}

type MapBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *MapBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// ModelProp implements types.ModelPropBuilder, build a property.
func (b *MapBuilder) ModelProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *MapBuilder) StructTag(key, value string) *MapBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *MapBuilder) Validator(v ...string) *MapBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *MapBuilder) Title(title string) *MapBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *MapBuilder) Description(desc string) *MapBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *MapBuilder) Example(example string) *MapBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *MapBuilder) Required() *MapBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to Map.
func (b *MapBuilder) GoType(i interface{}) *MapBuilder {
	b.goType(i, mapType)
	return b
}

type PBoolBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *PBoolBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *PBoolBuilder) StructTag(key, value string) *PBoolBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *PBoolBuilder) Validator(v ...string) *PBoolBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *PBoolBuilder) Title(title string) *PBoolBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *PBoolBuilder) Description(desc string) *PBoolBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *PBoolBuilder) Example(example string) *PBoolBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *PBoolBuilder) Required() *PBoolBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to PBool.
func (b *PBoolBuilder) GoType(i interface{}) *PBoolBuilder {
	b.goType(i, boolType)
	return b
}

type PStringBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *PStringBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *PStringBuilder) StructTag(key, value string) *PStringBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *PStringBuilder) Validator(v ...string) *PStringBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *PStringBuilder) Title(title string) *PStringBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *PStringBuilder) Description(desc string) *PStringBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *PStringBuilder) Example(example string) *PStringBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *PStringBuilder) Required() *PStringBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to PString.
func (b *PStringBuilder) GoType(i interface{}) *PStringBuilder {
	b.goType(i, stringType)
	return b
}

// MaximumLength set maximum length of string.
func (b *PStringBuilder) MaximumLength(max uint) *PStringBuilder {
	b.maximumLength(max)
	return b
}

// MinimumLength set maximum length of string.
func (b *PStringBuilder) MinimumLength(min uint) *PStringBuilder {
	b.minimumLength(min)
	return b
}

// Pattern to set value pattern with regular expression.
// As go-validator not support regular expression will (because some the performance reason),
// You should set a go-validator struct filed.
func (b *PStringBuilder) Pattern(regex string, validators ...string) *PStringBuilder {
	b.pattern(regex, validators...)
	return b
}

// Enum to set the valid values for string property.
func (b *PStringBuilder) Enum(values ...string) *PStringBuilder {
	b.enum(values...)
	return b
}

type PIntBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *PIntBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *PIntBuilder) StructTag(key, value string) *PIntBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *PIntBuilder) Validator(v ...string) *PIntBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *PIntBuilder) Title(title string) *PIntBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *PIntBuilder) Description(desc string) *PIntBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *PIntBuilder) Example(example string) *PIntBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *PIntBuilder) Required() *PIntBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to PInt.
func (b *PIntBuilder) GoType(i interface{}) *PIntBuilder {
	b.goType(i, intType)
	return b
}

// Maximum set the maximum value of this property.
func (b *PIntBuilder) Maximum(max int) *PIntBuilder {
	b.maximum(&max, nil)
	return b
}

// Minimum set the minimum value of this property.
func (b *PIntBuilder) Minimum(min int) *PIntBuilder {
	b.minimum(&min, nil)
	return b
}

type PFloatBuilder struct {
	builder
}

// ParameterProp implements types.ParameterPropBuilder, build a property.
func (b *PFloatBuilder) ParameterProp() *desc.Prop {
	return b.prop
}

// StructTag add a struct tag in this property.
func (b *PFloatBuilder) StructTag(key, value string) *PFloatBuilder {
	b.structTag(key, value)
	return b
}

// Validator add a go-validator tag in this property.
func (b *PFloatBuilder) Validator(v ...string) *PFloatBuilder {
	b.validator(v...)
	return b
}

// Title set title for this property. The title should be a single line text to introduce this propery.
func (b *PFloatBuilder) Title(title string) *PFloatBuilder {
	b.title(title)
	return b
}

// Description set description for this property.
// The description can be a multi line text to tell details of this propery.
func (b *PFloatBuilder) Description(desc string) *PFloatBuilder {
	b.description(desc)
	return b
}

// Example set example values for this property.
func (b *PFloatBuilder) Example(example string) *PFloatBuilder {
	b.example(example)
	return b
}

// Required mark this property is required
func (b *PFloatBuilder) Required() *PFloatBuilder {
	b.required()
	return b
}

// GoType to set override type of this property. The type must be convertible to PFloat.
func (b *PFloatBuilder) GoType(i interface{}) *PFloatBuilder {
	b.goType(i, floatType)
	return b
}

// Maximum set the maximum value of this property.
func (b *PFloatBuilder) Maximum(max float64) *PFloatBuilder {
	b.maximum(nil, &max)
	return b
}

// Minimum set the minimum value of this property.
func (b *PFloatBuilder) Minimum(min float64) *PFloatBuilder {
	b.minimum(nil, &min)
	return b
}
