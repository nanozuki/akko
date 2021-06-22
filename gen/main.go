package main

import (
	"fmt"
	"os"
	"text/template"
)

type BuilderParams struct {
	Types []BuilderParam
}

type BuilderParam struct {
	Type          string
	IsParameter   bool
	GoType        string
	InPath        bool
	InQuery       bool
	PType         string
	MaximumInt    bool
	MinimumInt    bool
	MaximumFloat  bool
	MinimumFloat  bool
	MaximumLength bool
	MinimumLength bool
	Pattern       bool
	Enum          bool
	MaximumItems  bool
	MinimumItems  bool
	Unique        bool
}

var defs = &BuilderParams{
	Types: []BuilderParam{
		{
			Type:    "Bool",
			GoType:  "bool",
			InQuery: true,
			PType:   "PBool",
		},
		{
			Type:          "String",
			GoType:        "string",
			InPath:        true,
			InQuery:       true,
			PType:         "PString",
			MaximumLength: true,
			MinimumLength: true,
			Pattern:       true,
			Enum:          true,
		},
		{
			Type:       "Int",
			GoType:     "int",
			InPath:     true,
			InQuery:    true,
			PType:      "PInt",
			MaximumInt: true,
			MinimumInt: true,
		},
		{
			Type:         "Float",
			GoType:       "float",
			InPath:       true,
			InQuery:      true,
			PType:        "PFloat",
			MaximumFloat: true,
			MinimumFloat: true,
		},
		{
			Type:         "Array",
			MaximumItems: true,
			MinimumItems: true,
			Unique:       true,
		},
		{
			Type: "Object",
		},
		{
			Type:   "Map",
			GoType: "map",
		},
	},
}

func (b *BuilderParams) build() {
	var pTypes []BuilderParam
	for _, t := range b.Types {
		if t.PType != "" {
			t.Type = t.PType
			t.IsParameter = true

			t.InPath = false
			t.InQuery = false
			t.PType = ""
			pTypes = append(pTypes, t)
		}
	}
	b.Types = append(b.Types, pTypes...)
}

func main() {
	tmpl := template.Must(template.ParseFiles("./gen/builder.tmpl"))
	filename := "./prop/typed_builders.go"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(fmt.Errorf("open file: %w", err))
	}
	defer file.Close()
	defs.build()
	if err := tmpl.Execute(file, defs); err != nil {
		panic(fmt.Errorf("exec template: %w", err))
	}
}
