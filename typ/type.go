package typ

import (
	"fmt"
	"reflect"

	"github.com/nanozuki/ononoki/desc"
	uuid "github.com/satori/go.uuid"
)

// ModelPropBuilder build model property, and model property can also build parameter property.
type ModelPropBuilder interface {
	ParameterProp() *desc.Prop
	ModelProp() *desc.Prop
}

// ParameterPropBuilder build property only can be use in parameter.
type ParameterPropBuilder interface {
	ParameterProp() *desc.Prop
}

// Bool returns a TypeInfo with type bool.
func Bool() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeBool,
		Ident: "bool",
	}
}

// String returns a TypeInfo with type string.
func String() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeString,
		Ident: "string",
	}
}

// Int returns a TypeInfo with type int.
func Int() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeInt,
		Ident: "int32",
	}
}

// Float returns a TypeInfo with type float.
func Float() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeFloat,
		Ident: "float64",
	}
}

// Array returns a TypeInfo with type array.
func Array(itemTyp *desc.TypeInfo) *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeArray,
		Ident: "[]" + itemTyp.Ident,
		Item:  itemTyp,
	}
}

// Object to create a object type with giving properties.
func Object(name string, builders ...ModelPropBuilder) *desc.TypeInfo {
	var props []*desc.Prop
	for _, build := range builders {
		props = append(props, build.ModelProp())
	}
	return &desc.TypeInfo{
		Type:     desc.TypeObject,
		Ident:    string(desc.NewName(name)),
		Props:    props,
		ObjectID: uuid.NewV4().String(),
	}
}

// GoType returns a TypeInfo extract from golang value.
func GoType(i interface{}) *desc.TypeInfo {
	t := reflect.TypeOf(i)
	tv := indirect(t)
	ti := &desc.TypeInfo{
		Ident:   t.String(),
		PkgPath: tv.PkgPath(),
	}
	switch tv.Kind() {
	case reflect.Bool:
		ti.Type = desc.TypeBool
	case reflect.String:
		ti.Type = desc.TypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		ti.Type = desc.TypeInt
	case reflect.Float32, reflect.Float64:
		ti.Type = desc.TypeFloat
	case reflect.Array, reflect.Slice:
		ti.Type = desc.TypeArray
	case reflect.Struct:
		ti.Type = desc.TypeObject
	case reflect.Map:
		ti.Type = desc.TypeMap
	default:
		panic(fmt.Errorf("not support %q", tv))
	}
	return ti
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// Map returns a TypeInfo with type map[string]interface{}.
func Map() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeMap,
		Ident: "map[string]interface{}",
	}
}
