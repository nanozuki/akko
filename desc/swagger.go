package desc

import (
	"fmt"
	"strings"
)

// SwaggerProp present properties schema in swagger.
type SwaggerProp struct {
	Title         string   `json:"title,omitempty"`
	Description   string   `json:"description,omitempty"`
	Example       string   `json:"example,omitempty"`
	Required      bool     `json:"required,omitempty"`
	In            Location `json:"in,omitempty"`
	Maximum       string   `json:"maximum,omitempty"`        // for number
	Minimum       string   `json:"minimum,omitempty"`        // for number
	MaximumLength uint     `json:"maximum_length,omitempty"` // for string
	MinimumLength uint     `json:"minimum_length,omitempty"` // for string
	Pattern       string   `json:"pattern,omitempty"`        // for string
	Enum          []string `json:"enum,omitempty"`           // for string
	MaximumItems  uint     `json:"maximum_items,omitempty"`  // for array
	MinimumItems  uint     `json:"minimum_items,omitempty"`  // for array
	Unique        bool     `json:"unique,omitempty"`         // for array
}

const (
	pLinePre = "\t// "
	oLinePre = "// "
)

// ToComment build go-swagger field comment.
func (sp SwaggerProp) ToComment() string {
	var strs []string
	if sp.Title != "" {
		// title must be single line text
		strs = append(strs, pLinePre+strings.ReplaceAll(sp.Title, "\n", " "))
	}
	if sp.Description != "" {
		// title can be multi line text
		desc := strings.Split(sp.Description, "\n")
		if len(strs) != 0 {
			strs = append(strs, pLinePre)
		}
		for _, d := range desc {
			strs = append(strs, pLinePre+d)
		}
	}
	{
		var exStrs []string
		if sp.Example != "" {
			exStrs = append(exStrs, pLinePre+"example: ", sp.Example)
		}
		if sp.Required {
			exStrs = append(exStrs, pLinePre+"required: true")
		}
		if sp.In != Unset {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("in: %v", sp.In))
		}
		if sp.Maximum != "" {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("maximum: %v", sp.Maximum))
		}
		if sp.Minimum != "" {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("minimum: %v", sp.Minimum))
		}
		if sp.MaximumLength != 0 {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("maximumLength: %v", sp.MaximumLength))
		}
		if sp.MinimumLength != 0 {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("minimumLength: %v", sp.MinimumLength))
		}
		if sp.Pattern != "" {
			exStrs = append(exStrs, pLinePre+"pattern: "+sp.Pattern)
		}
		if len(sp.Enum) != 0 {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("[%v]", strings.Join(sp.Enum, " ")))
		}
		if sp.MaximumItems != 0 {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("maximumItems: %v", sp.MaximumItems))
		}
		if sp.MinimumItems != 0 {
			exStrs = append(exStrs, pLinePre+fmt.Sprintf("minimumItems: %v", sp.MinimumItems))
		}
		if sp.Unique {
			exStrs = append(exStrs, pLinePre+"unique: true")
		}
		if len(exStrs) != 0 {
			if len(strs) != 0 {
				strs = append(strs, pLinePre)
			}
			strs = append(strs, exStrs...)
		}
	}
	return strings.Join(strs, "\n")
}

// SwaggerObject present attributes for object in swagger.
type SwaggerObject struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Parameter   string `json:"parameter,omitempty"`
}

// ToComment build go-swagger object comment.
func (sp SwaggerObject) ToComment() string {
	var strs []string
	if sp.Title != "" {
		// title must be single line text
		strs = append(strs, oLinePre+strings.ReplaceAll(sp.Title, "\n", " "))
	}
	if sp.Description != "" {
		// title can be multi line text
		desc := strings.Split(sp.Description, "\n")
		if len(strs) != 0 {
			strs = append(strs, oLinePre)
		}
		for _, d := range desc {
			strs = append(strs, oLinePre+d)
		}
	}
	if len(strs) != 0 {
		strs = append(strs, oLinePre)
	}
	if sp.Parameter != "" {
		strs = append(strs, oLinePre+"swagger:parameters "+sp.Parameter)
	} else {
		strs = append(strs, oLinePre+"swagger:model")
	}
	return strings.Join(strs, "\n")
}

// Location for properties in parameter.
type Location string

const (
	Unset   Location = ""
	InBody  Location = "body"
	InPath  Location = "path"
	InQuery Location = "query"
)
