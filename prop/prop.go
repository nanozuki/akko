package prop

import (
	"github.com/nanozuki/ononoki/desc"
	"github.com/nanozuki/ononoki/typ"
)

func Bool(name string) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Bool(),
	}
}

/* TODO
func Enum() *desc.Prop
*/

func String(name string) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.String(),
	}
}

func Int(name string) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Int(),
	}
}

func Float(name string) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Float(),
	}
}

func Array(name string, itemTyp *desc.TypeInfo) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Array(itemTyp),
	}
}

func Object(name string, p ...*desc.Prop) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Object(name, p...),
	}
}

func GoType(name string, i interface{}) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.GoType(i),
	}
}

func Map(name string) *desc.Prop {
	return &desc.Prop{
		Name: name,
		Type: typ.Map(),
	}
}
