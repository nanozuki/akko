package typ

import (
	"reflect"

	"github.com/nanozuki/ononoki/desc"
)

func Bool() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeBool,
		Ident: "bool",
	}
}

/* TODO
func Enum() *desc.TypeInfo {}
*/

func String() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeString,
		Ident: "string",
	}
}

func Int() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeInt,
		Ident: "int32",
	}
}

func Float() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeFloat,
		Ident: "float64",
	}
}

func Array(itemTyp *desc.TypeInfo) *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeArray,
		Ident: "[]" + itemTyp.Ident,
		Item:  itemTyp,
	}
}

func Object(name string, props ...*desc.Prop) *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeObject,
		Ident: name,
		Props: props,
	}
}

func GoType(i interface{}) *desc.TypeInfo {
	t := reflect.TypeOf(i)
	tv := indirect(t)
	return &desc.TypeInfo{
		Type:    desc.TypeGoType,
		Ident:   t.String(),
		PkgPath: tv.PkgPath(),
	}
}

// indirect returns the type at the end of indirection.
func indirect(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func Map() *desc.TypeInfo {
	return &desc.TypeInfo{
		Type:  desc.TypeMap,
		Ident: "map[string]interface{}",
	}
}
