package ononoki

import (
	"fmt"
	"net/http"
	p "path"

	"github.com/nanozuki/ononoki/desc"
	"github.com/nanozuki/ononoki/prop"
	"github.com/nanozuki/ononoki/typ"
)

// Make sure operation name and endpoint is unique.
var (
	opNames   = map[desc.Name]bool{}
	endpoints = map[string]bool{}
)

// Schema definition a service.
type Schema struct {
	name        string
	middlewares []Middleware
	builders    []*RouteBuilder
	routes      []*Route
}

// New create schema with name.
func New(name string) *Schema {
	return &Schema{
		name: name,
	}
}

// Middleware is a string placeholder for middleware.
type Middleware string

// Example for default middleware
const (
	LogMiddleware          Middleware = "Log"
	AdminAuthMiddleware    Middleware = "AdminAuth"
	CustomerAuthMiddleware Middleware = "CustomerAuth"
	RequestIDMiddleware    Middleware = "RequestID"
)

// Use adds middleware to the chain which is run after router.
func (s *Schema) Use(middleware ...Middleware) {
	s.middlewares = append(s.middlewares, middleware...)
}

// Group creates a new sub-group with prefix and optional sub-group-level middleware.
func (s *Schema) Group(prefix string, m ...Middleware) *Group {
	middlewares := make([]Middleware, 0, len(s.middlewares)+len(m))
	middlewares = append(middlewares, s.middlewares...)
	middlewares = append(middlewares, m...)
	return &Group{
		prefix:      prefix,
		middlewares: middlewares,
		tags:        nil,
		schema:      s,
	}
}

// Tag creates a new sub-group with tags.
func (s *Schema) Tag(tag ...string) *Group {
	middlewares := make([]Middleware, 0, len(s.middlewares))
	middlewares = append(middlewares, s.middlewares...)
	return &Group{
		prefix:      "/",
		middlewares: middlewares,
		tags:        tag,
		schema:      s,
	}
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (s *Schema) Add(method, name, path string) *RouteBuilder {
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
	default:
		panic("invalid method")
	}
	middlewares := make([]Middleware, 0, len(s.middlewares))
	middlewares = append(middlewares, s.middlewares...)
	route := s.add(name, method, p.Join("/", path), []string{}, middlewares...)
	return route
}

func (s *Schema) add(name, method, path string, tags []string, m ...Middleware) *RouteBuilder {
	opName := desc.NewName(name)
	if endpoints[method+path] {
		panic(fmt.Errorf("endpoint '%s %s' is duplicated", method, path))
	}
	if opNames[opName] {
		panic(fmt.Errorf("operation name '%s' is duplicated", name))
	}
	route := &RouteBuilder{
		r: &Route{
			Name:        opName,
			Tags:        tags,
			URL:         path,
			Method:      method,
			Middlewares: m,
		},
	}
	s.builders = append(s.builders, route)
	endpoints[method+path] = true
	opNames[opName] = true
	return route
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (s *Schema) GET(name, path string) *RouteBuilder {
	return s.Add(http.MethodGet, name, path)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (s *Schema) POST(name, path string) *RouteBuilder {
	return s.Add(http.MethodPost, name, path)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (s *Schema) PUT(name, path string) *RouteBuilder {
	return s.Add(http.MethodPut, name, path)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (s *Schema) DELETE(name, path string) *RouteBuilder {
	return s.Add(http.MethodDelete, name, path)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (s *Schema) HEAD(name, path string) *RouteBuilder {
	return s.Add(http.MethodHead, name, path)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (s *Schema) OPTIONS(name, path string) *RouteBuilder {
	return s.Add(http.MethodOptions, name, path)
}

// Name returns the name of schema
func (s Schema) Name() string {
	return s.name
}

// Routes returns all routes in schema.
func (s Schema) Routes() []*Route {
	return s.routes
}

// RouteBuilder is a helper struct to build Route.
type RouteBuilder struct {
	r *Route
}

// Request definition request object of route.
func (r *RouteBuilder) Request(props ...typ.ParameterPropBuilder) *RouteBuilder {
	r.r.Req = &desc.TypeInfo{
		Type:  desc.TypeObject,
		Ident: string(r.r.Name) + "Req",
	}
	var pProps []*desc.Prop
	var bodyProps []*desc.Prop
	for _, p := range props {
		prop := p.ParameterProp()
		switch prop.Swagger.In {
		case desc.InPath, desc.InQuery:
			pProps = append(pProps, prop)
		default:
			bodyProps = append(bodyProps, prop)
		}
	}
	switch {
	case len(pProps) == 0 && len(bodyProps) == 0:
	case len(pProps) != 0 && len(bodyProps) == 0:
		r.r.Req.Props = pProps
	case len(pProps) == 0 && len(bodyProps) != 0:
		r.r.Req.Props = bodyProps
	case len(pProps) != 0 && len(bodyProps) != 0:
		r.r.Req.Props = pProps
		body := prop.Object("body", typ.Object(string(r.r.Name)+"Req"+"Body")).ModelProp()
		body.TypeInfo.Props = bodyProps
		body.Swagger.In = desc.InBody
		r.r.Req.Props = append(r.r.Req.Props, body)
	}
	r.r.ReqHasBody = len(pProps) != 0 && len(bodyProps) != 0
	if len(pProps) != 0 {
		r.r.Req.Swagger.Parameter = string(r.r.Name)
	}
	return r
}

// Response definition response object of route.
func (r *RouteBuilder) Response(props ...typ.ModelPropBuilder) *RouteBuilder {
	r.r.Res = typ.Object(string(r.r.Name)+"Res", props...)
	return r
}

// Use adds middleware to the chain which is run after router.
func (r *RouteBuilder) Use(m ...Middleware) *RouteBuilder {
	r.r.Middlewares = append(r.r.Middlewares, m...)
	return r
}

// Description to add description for route.
func (r *RouteBuilder) Description(desc string) *RouteBuilder {
	r.r.Description = desc
	return r
}

// Route contains a handler and information for matching against requests.
type Route struct {
	Name        desc.Name      `json:"name"`
	Description string         `json:"description"`
	Tags        []string       `json:"tags"`
	URL         string         `json:"url"`
	Method      string         `json:"method"`
	Req         *desc.TypeInfo `json:"req,omitempty"`
	Res         *desc.TypeInfo `json:"res,omitempty"`
	ReqHasBody  bool           `json:"req_has_body,omitempty"`
	Middlewares []Middleware   `json:"middlewares"`
}
