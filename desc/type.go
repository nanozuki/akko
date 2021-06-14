package desc

// A Type represents a field type.
type Type uint8

// List of field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeEnum
	TypeString
	TypeInt
	TypeFloat

	TypeArray
	TypeObject
	TypeGoType
	TypeMap
)

// TypeInfo holds the information regarding types in dto.
type TypeInfo struct {
	Type    Type
	Ident   string // how to use this type in code
	PkgPath string // import path

	Item  *TypeInfo
	Props []*Prop
}
