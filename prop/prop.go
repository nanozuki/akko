//
package prop

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/nanozuki/ononoki/desc"
	"github.com/nanozuki/ononoki/typ"
)

/*
methods,types -> Bool String Int Float Array Object Map
    GoType        o     o     o    o                 o
    Title         o     o     o    o     o     o     o
    Desc          o     o     o    o     o     o     o
    Example       o     o     o    o     o     o     o
    Required      o     o     o    o     o     o     o
    InPath              o     o    o
    InQuery       o     o     o    o
    Maximum                   o    o
    Minimum                   o    o
 MaximumLength          o
 MinimumLength          o
    Pattern             o
    Enum                o
 MaximumItems                            o
 MinimumItems                            o
    Unique                               o
*/

// builder is the core of all typed builders, includes all logics for all types.
type builder struct {
	prop    *desc.Prop
	rawName string
}

func newBuilder(name string, typeInfo *desc.TypeInfo) builder {
	b := builder{
		prop:    &desc.Prop{Name: desc.NewName(name), TypeInfo: typeInfo},
		rawName: name,
	}
	b.prop.StructTags.Add("json", name)
	return b
}

func (b *builder) goType(i interface{}, expectType reflect.Type) {
	t := reflect.TypeOf(i)
	if t.Kind() == expectType.Kind() && t.ConvertibleTo(expectType) {
		b.prop.TypeInfo = typ.GoType(i)
	} else {
		panic(fmt.Errorf("must be %q", expectType))
	}
}

func (b *builder) structTag(key, value string) {
	b.prop.StructTags.Add(key, value)
}

func (b *builder) validator(v ...string) {
	b.prop.StructTags.Validator(v...)
}

func (b *builder) title(title string) {
	b.prop.Swagger.Title = title
}

func (b *builder) description(desc string) {
	b.prop.Swagger.Description = desc
}

func (b *builder) example(example string) {
	b.prop.Swagger.Example = example
}

func (b *builder) required() {
	b.prop.Swagger.Required = true
	b.prop.StructTags.Validator("required")
}

func (b *builder) inPath() {
	b.prop.Swagger.In = desc.InPath
	b.prop.StructTags.Add("param", b.rawName)

}

func (b *builder) inQuery() {
	b.prop.Swagger.In = desc.InQuery
	b.prop.StructTags.Add("query", b.rawName)
}

func (b *builder) maximum(i *int, f *float64) {
	max := ""
	if i != nil {
		max = fmt.Sprint(*i)
	} else if f != nil {
		max = fmt.Sprint(*f)
	}
	b.prop.Swagger.Maximum = max
	b.prop.StructTags.Validator(fmt.Sprintf("max=%v", max))
}

func (b *builder) minimum(i *int, f *float64) {
	min := ""
	if i != nil {
		min = fmt.Sprint(*i)
	} else if f != nil {
		min = fmt.Sprint(*f)
	}
	b.prop.Swagger.Minimum = min
	b.prop.StructTags.Validator(fmt.Sprintf("min=%v", min))
}

func (b *builder) maximumLength(max uint) {
	b.prop.Swagger.MaximumLength = max
	b.prop.StructTags.Validator(fmt.Sprintf("max=%v", max))
}

func (b *builder) minimumLength(min uint) {
	b.prop.Swagger.MinimumLength = min
	b.prop.StructTags.Validator(fmt.Sprintf("min=%v", min))
}

func (b *builder) pattern(regex string, validators ...string) {
	b.prop.Swagger.Pattern = regex
	b.prop.StructTags.Validator(validators...)
}

func (b *builder) enum(values ...string) {
	b.prop.Swagger.Enum = values
	var vs []string
	for _, v := range values {
		if strings.Contains(v, " ") {
			vs = append(vs, "'"+v+"'")
		} else {
			vs = append(vs, v)
		}
	}
	b.prop.StructTags.Validator(fmt.Sprintf("oneOf=%s", strings.Join(vs, " ")))
}

func (b *builder) maximumItems(max uint) {
	b.prop.Swagger.MaximumItems = max
	b.prop.StructTags.Validator(fmt.Sprintf("max=%v", max))
}

func (b *builder) minimumItems(min uint) {
	b.prop.Swagger.MinimumItems = min
	b.prop.StructTags.Validator(fmt.Sprintf("min=%v", min))
}

func (b *builder) unique() {
	b.prop.Swagger.Unique = true
	b.prop.StructTags.Validator("unique")
}

var (
	boolType   = reflect.TypeOf(false)
	stringType = reflect.TypeOf("")
	intType    = reflect.TypeOf(int(0))
	floatType  = reflect.TypeOf(float64(0.0))
	mapType    = reflect.TypeOf(map[string]interface{}{})
)

// Bool returns a property builder with type bool.
func Bool(name string) *BoolBuilder {
	return &BoolBuilder{newBuilder(name, typ.Bool())}
}

// String returns a property builder with type string.
func String(name string) *StringBuilder {
	return &StringBuilder{newBuilder(name, typ.String())}
}

// Int returns a property builder with type int.
func Int(name string) *IntBuilder {
	return &IntBuilder{newBuilder(name, typ.Int())}
}

// Float returns a property builder with type int.
func Float(name string) *FloatBuilder {
	return &FloatBuilder{newBuilder(name, typ.Float())}
}

// Array returns a property builder with type array. Should specify the item type.
func Array(name string, itemType *desc.TypeInfo) *ArrayBuilder {
	return &ArrayBuilder{newBuilder(name, typ.Array(itemType))}
}

// Object returns a property builder with object, should specify the object type.
// If type info is not a object, this function will panic.
func Object(propName string, typeInfo *desc.TypeInfo) *ObjectBuilder {
	if typeInfo.Type != desc.TypeObject {
		panic(errors.New("must be object"))
	}
	return &ObjectBuilder{newBuilder(propName, typeInfo)}
}

// GoType returns a property builder with the type of giving go value.
func GoType(propName string, i interface{}) *ObjectBuilder {
	return &ObjectBuilder{newBuilder(propName, typ.GoType(i))}
}

// Map returns a property builder with the type of 'map[string]interface{}',
// It's used by dynamic json object.
func Map(name string) *MapBuilder {
	return &MapBuilder{newBuilder(name, typ.Map())}
}
