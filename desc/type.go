package desc

import "fmt"

// A Type represents a field type.
type Type uint8

// List of field types.
const (
	TypeInvalid Type = iota
	TypeBool
	TypeString
	TypeInt
	TypeFloat

	TypeArray
	TypeObject // for open-api object
	TypeMap    // map for json, map[string]interface{}
)

// TypeInfo holds the information regarding types in dto.
type TypeInfo struct {
	Type    Type          `json:"type,omitempty"`
	Ident   string        `json:"ident,omitempty"`    // how to use this type in code
	PkgPath string        `json:"pkg_path,omitempty"` // import path
	Swagger SwaggerObject `json:"swagger,omitempty"`
	// Related values
	Item     *TypeInfo `json:"item,omitempty"`      // for array item type
	Values   []string  `json:"values,omitempty"`    // for enum's values
	Props    []*Prop   `json:"props,omitempty"`     // for object's properties
	ObjectID string    `json:"object_id,omitempty"` // for object's generator
}

// Title to set title for type.
func (ti *TypeInfo) Title(title string) *TypeInfo {
	if ti.Type != TypeObject {
		fmt.Println("[Warning!!!]: title is only useful for object")
	}
	ti.Swagger.Title = title
	return ti
}

// Description to set description for type.
func (ti *TypeInfo) Description(desc string) *TypeInfo {
	if ti.Type != TypeObject {
		fmt.Println("[Warning!!!]: description is only useful for object")
	}
	ti.Swagger.Description = desc
	return ti
}
