package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"

	"golang.org/x/tools/imports"
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
	defs.build()
	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, defs); err != nil {
		panic(fmt.Errorf("exec template: %w", err))
	}
	src, err := imports.Process(filename, nil, nil)
	if err != nil {
		panic(fmt.Errorf("format file %s: %w", filename, err))
	}
	if err := ioutil.WriteFile(filename, src, 0644); err != nil {
		panic(fmt.Errorf("write file %s: %w", filename, err))
	}

}
