package ononoki

import (
	"net/http"
	p "path"
)

// Group is a set of sub-routes for a specified route.
type Group struct {
	prefix      string
	middlewares []Middleware
	tags        []string
	schema      *Schema
}

// Use adds middleware to the chain which is run after router.
func (g *Group) Use(middleware ...Middleware) {
	g.middlewares = append(g.middlewares, middleware...)
}

// Group creates a new sub-group with prefix and optional sub-group-level middleware.
func (g *Group) Group(prefix string, m ...Middleware) *Group {
	middlewares := make([]Middleware, 0, len(g.middlewares)+len(m))
	middlewares = append(middlewares, g.middlewares...)
	middlewares = append(middlewares, m...)
	return &Group{
		prefix:      p.Join(g.prefix, prefix),
		middlewares: middlewares,
		tags:        g.tags,
		schema:      g.schema,
	}
}

// Tag creates a new sub-group with tags.
func (g *Group) Tag(tag ...string) *Group {
	middlewares := make([]Middleware, 0, len(g.middlewares))
	middlewares = append(middlewares, g.middlewares...)
	return &Group{
		prefix:      g.prefix,
		middlewares: middlewares,
		tags:        tag,
		schema:      g.schema,
	}
}

// Add registers a new route for an HTTP method and path with matching handler
// in the router with optional route-level middleware.
func (g *Group) Add(method, name, path string) *RouteBuilder {
	middlewares := make([]Middleware, 0, len(g.middlewares))
	middlewares = append(middlewares, g.middlewares...)
	return g.schema.add(name, method, p.Join(g.prefix, path), g.tags, middlewares...)
}

// GET registers a new GET route for a path with matching handler in the router
// with optional route-level middleware.
func (g *Group) GET(name, path string) *RouteBuilder {
	return g.Add(http.MethodGet, name, path)
}

// POST registers a new POST route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) POST(name, path string) *RouteBuilder {
	return g.Add(http.MethodPost, name, path)
}

// PUT registers a new PUT route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) PUT(name, path string) *RouteBuilder {
	return g.Add(http.MethodPut, name, path)
}

// DELETE registers a new DELETE route for a path with matching handler in the router
// with optional route-level middleware.
func (g *Group) DELETE(name, path string) *RouteBuilder {
	return g.Add(http.MethodDelete, name, path)
}

// HEAD registers a new HEAD route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) HEAD(name, path string) *RouteBuilder {
	return g.Add(http.MethodHead, name, path)
}

// OPTIONS registers a new OPTIONS route for a path with matching handler in the
// router with optional route-level middleware.
func (g *Group) OPTIONS(name, path string) *RouteBuilder {
	return g.Add(http.MethodOptions, name, path)
}
