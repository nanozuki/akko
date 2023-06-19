package akko

import (
	"errors"
	"fmt"
	"go/ast"
	"net/http"
	"reflect"
	"strings"
)

type TypeInfo struct {
	Name      string
	PkgPath   string
	Indentity string
	Kind      reflect.Kind
}

func NewTypeInfo(t reflect.Type) TypeInfo {
	ti := TypeInfo{
		Name:      t.Name(),
		PkgPath:   t.PkgPath(),
		Indentity: t.String(),
		Kind:      t.Kind(),
	}
	if ti.Kind == reflect.Ptr {
		ti.PkgPath = t.Elem().PkgPath()
	}
	return ti
}

type ServiceInfo struct {
	Typ     TypeInfo
	Methods []MethodInfo
}

func NewServiceInfo(service any) (*ServiceInfo, error) {
	t := reflect.TypeOf(service)
	si := &ServiceInfo{
		Typ: NewTypeInfo(t),
	}
	si.Methods = make([]MethodInfo, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		if err := (&si.Methods[i]).ReadType(t, i); err != nil {
			return nil, err
		}
	}
	return si, nil
}

// FillByAst fill service info by parsing directory
// Find the comment of service methods, and fill to service methods info
// Find the argument name of service methods, and fill to service methods info
func (s *ServiceInfo) FillByAst(pkg *ast.Package) error {
	return nil
}

type MethodInfo struct {
	Name       string
	Inputs     []ParamInfo
	Returns    []ParamInfo
	HTTPMethod string
	URL        string
	Process    []Process
}

func (m *MethodInfo) DeclString() string {
	var b strings.Builder
	b.WriteString(m.Name)
	b.WriteString("(")
	for i, in := range m.Inputs {
		b.WriteString(in.String())
		if i != len(m.Inputs)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(") -> (")
	for i, out := range m.Returns {
		b.WriteString(out.String())
		if i != len(m.Returns)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(")")
	return b.String()
}

var ErrInvalidComment = errors.New("invalid method comment")

func IsValidHTTPMethods(m string) bool {
	switch m {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
		return true
	default:
		return false
	}
}

// ReadType read i-th method of receiver to fill Name, Inputs and Returns
func (m *MethodInfo) ReadType(receiver reflect.Type, i int) error {
	method := receiver.Method(i)
	m.Name = method.Name
	m.Inputs = make([]ParamInfo, 0, method.Type.NumIn())
	m.Returns = make([]ParamInfo, 0, method.Type.NumOut())
	for i := 1; i < method.Type.NumIn(); i++ {
		m.Inputs = append(m.Inputs, ParamInfo{
			// Name: method.Type.In(i).Name(), TODO: get param name from comment
			Typ: NewTypeInfo(method.Type.In(i)),
		})
	}
	for i := 0; i < method.Type.NumOut(); i++ {
		m.Returns = append(m.Returns, ParamInfo{
			// Name: method.Type.Out(i).Name(), TODO: get param name from comment
			Typ: NewTypeInfo(method.Type.Out(i)),
		})
	}
	return nil
}

// ParseComment parse comment like "[PATCH /user, body|json->patch, return|json->body]" to HTTPMethod, URL and Process
func (m *MethodInfo) ParseComment(comment string) error {
	const minCommentLen = 2 + 3 + 1 + 1 // 2: square brackets, 3: min HTTP method length, 1: space, 1: min URL length
	if len(comment) < minCommentLen || comment[0] != '[' || comment[len(comment)-1] != ']' {
		return ErrInvalidComment
	}
	comment = comment[1 : len(comment)-1]
	parts := strings.Split(comment, ",")
	{
		// parse parts[0] to HTTPMethod and URL
		pp := strings.Split(strings.TrimSpace(parts[0]), " ")
		if len(pp) != 2 || !IsValidHTTPMethods(pp[0]) {
			return ErrInvalidComment
		}
		m.HTTPMethod = pp[0]
		m.URL = pp[1]
	}
	{
		// parse parts[1:] to Processes
		for _, pp := range parts[1:] {
			p, err := ParseProcess(pp)
			if err != nil {
				return err
			}
			m.Process = append(m.Process, p)
		}
	}
	return nil
}

type ParamInfo struct {
	Name string
	Typ  TypeInfo
}

func (i ParamInfo) String() string {
	if i.Name == "" {
		return i.Typ.Indentity
	}
	return fmt.Sprintf("%s %s", i.Name, i.Typ.Indentity)
}

type Process struct {
	Input    ProcessIn
	Methods  []string
	Pipeline []string
	Output   ProcessOut
}

var ErrInvalidProcess = errors.New("invalid process")

// ParseProcess parse string like "params.id->id" or "return|json->body" to Process
func ParseProcess(s string) (Process, error) {
	p := Process{}
	parts := strings.Split(s, "->")
	if len(parts) != 2 {
		return p, ErrInvalidProcess
	}
	p.Output = ProcessOut(parts[1])
	if !p.Output.IsValid() {
		return p, ErrInvalidProcess
	}
	parts = strings.Split(parts[0], "|")
	if len(parts) > 1 {
		p.Pipeline = parts[1:]
	}
	parts = strings.Split(parts[0], ".")
	if len(parts) > 1 {
		p.Methods = parts[1:]
	}
	p.Input = ProcessIn(parts[0])
	if !p.Input.IsValid() {
		return p, ErrInvalidProcess
	}
	return p, nil
}

func (p Process) String() string {
	var b strings.Builder
	b.WriteString(string(p.Input))
	for _, m := range p.Methods {
		b.WriteString(".")
		b.WriteString(m)
	}
	for _, p := range p.Pipeline {
		b.WriteString("|")
		b.WriteString(p)
	}
	b.WriteString("->")
	b.WriteString(string(p.Output))
	return b.String()
}

type ProcessIn string

const (
	HeaderProcessIn ProcessIn = "header"
	ParamsProcessIn ProcessIn = "params"
	QueryProcessIn  ProcessIn = "query"
	BodyProcessIn   ProcessIn = "body"
)

func (p ProcessIn) IsValid() bool {
	switch p {
	case HeaderProcessIn, ParamsProcessIn, QueryProcessIn, BodyProcessIn:
		return true
	default:
		return false
	}
}

type ProcessOut string

const (
	HeaderProcessOut ProcessOut = "header"
	BodyProcessOut   ProcessOut = "body"
)

func (p ProcessOut) IsValid() bool {
	switch p {
	case HeaderProcessOut, BodyProcessOut:
		return true
	default:
		return false
	}
}
