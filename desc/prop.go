package desc

import (
	"fmt"
	"strings"
)

// Prop descripte a property in open-api,
// it will be a field in request or response or object.
type Prop struct {
	Name       Name        `json:"name,omitempty"`
	TypeInfo   *TypeInfo   `json:"type_info,omitempty"`
	StructTags StructTags  `json:"struct_tags,omitempty"`
	Swagger    SwaggerProp `json:"swagger,omitempty"`
}

// StructTags help to build structTag in struct field.
type StructTags struct {
	Tags []*structTag `json:"tags,omitempty"`
}

type structTag struct {
	Key   string   `json:"key,omitempty"`
	Value []string `json:"value,omitempty"`
}

// Set a struct tag, if tag already exisit, it will replace the value.
func (ts *StructTags) Set(key, value string) {
	for _, t := range ts.Tags {
		if t.Key == key {
			t.Value = []string{value}
			return
		}
	}
	ts.Tags = append(ts.Tags, &structTag{Key: key, Value: []string{value}})
}

// Add a struct tag, if tag already exisit, it will append new value at end.
func (ts *StructTags) Add(key, value string) {
	for _, t := range ts.Tags {
		if t.Key == key {
			for _, v := range t.Value {
				if v == value {
					return
				}
			}
			t.Value = append(t.Value, value)
			return
		}
	}
	ts.Tags = append(ts.Tags, &structTag{Key: key, Value: []string{value}})
}

// Get a tag value, if tag is not exist, return empty string.
func (ts *StructTags) Get(key string) string {
	for _, t := range ts.Tags {
		if t.Key == key {
			return strings.Join(t.Value, ",")
		}
	}
	return ""
}

// String return the structTag format in field.
func (ts *StructTags) String() string {
	var strs []string
	for _, tag := range ts.Tags {
		strs = append(strs, fmt.Sprintf("%s:\"%s\"", tag.Key, strings.Join(tag.Value, ",")))
	}
	return fmt.Sprintf("`%s`", strings.Join(strs, " "))
}

// Validator to help adding go-validator's validator in structTag.
func (ts *StructTags) Validator(v ...string) {
	for _, vv := range v {
		ts.Add("validator", vv)
	}
}
