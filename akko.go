package akko

import (
	"encoding/json"
	"fmt"
	"go/parser"
	"go/token"
	"log"
	"net/http"
	"reflect"
	"strings"
)

type RouteGroup struct {
	Domain    string
	Base      string
	Service   *ServiceInfo
	Providers map[string]*ProviderInfo
}

func (g *RouteGroup) String() string {
	domain := g.Domain
	if domain == "" {
		domain = "*"
	}
	strs := []string{
		fmt.Sprintf("Service: %s (%s)", g.Service.Typ.Indentity, g.Service.Typ.PkgPath),
		fmt.Sprintf("\tDomain: %s", domain),
		fmt.Sprintf("\tBase: %s", g.Base),
		"\tMethods:",
	}
	for _, m := range g.Service.Methods {
		strs = append(strs, fmt.Sprintf("\t\t%s", m.DeclString()))
	}
	strs = append(strs, "\tProviders:")
	for _, provider := range g.Providers {
		strs = append(strs, fmt.Sprintf("\t\t%s.%s", provider.Func.PkgPath(), provider.Func.Name()))
	}
	return strings.Join(strs, "\n")
}

type MountOption func(*RouteGroup)

type Provider[T any] func(*http.Request) (T, error)

type ProviderInfo struct {
	Func    reflect.Type
	Product reflect.Type
}

func NewProviderInfo[T any](p Provider[T]) ProviderInfo {
	var t T
	pi := ProviderInfo{
		Func:    reflect.TypeOf(p),
		Product: reflect.TypeOf(t),
	}
	if pi.Product.Kind() == reflect.Ptr {
		pi.Product = pi.Product.Elem()
	}
	return pi
}

func (i *ProviderInfo) Provide() string {
	if i.Product.PkgPath() == "" {
		return i.Product.Name()
	}
	return fmt.Sprintf("%s.%s", i.Product.PkgPath(), i.Product.Name())
}

func WithProvider[T any](p Provider[T]) MountOption {
	return func(g *RouteGroup) {
		if g.Providers == nil {
			g.Providers = make(map[string]*ProviderInfo)
		}
		pi := NewProviderInfo(p)
		g.Providers[pi.Provide()] = &pi
		fmt.Printf("get provider for %s\n", pi.Provide())
	}
}

type RequestHandler func(*http.Request) (*http.Request, error)

type ReturnHandler func(ret any, w http.ResponseWriter)

type ErrorHandler func(err error, w http.ResponseWriter)

func OnRequest(h RequestHandler) MountOption {
	return func(*RouteGroup) {}
}

func OnResponse(h ReturnHandler) MountOption {
	return func(*RouteGroup) {}
}

func OnError(h ErrorHandler) MountOption {
	return func(*RouteGroup) {}
}

func jsonResponse(ret any, w http.ResponseWriter) {
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var ReturnHandlers = map[string]ReturnHandler{
	"json": jsonResponse,
}

var RouteGroups []*RouteGroup

func Mount[Srv any](baseRoute string, service Srv, options ...MountOption) *RouteGroup {
	srv, err := NewServiceInfo(service)
	if err != nil {
		log.Fatalf("parse service %s failed: %s", srv.Typ.Indentity, err)
	}
	g := RouteGroup{
		Base:    baseRoute,
		Service: srv,
	}
	for _, opt := range options {
		opt(&g)
	}
	RouteGroups = append(RouteGroups, &g)
	return &g
}

func ParseDirectory(dir string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("parse directory %s failed: %s", dir, err)
	}
	for pkg, pkgAst := range pkgs {
		fmt.Printf("parse ast for package: %s\n", pkg)
		for _, rg := range RouteGroups {
			if rg.Service.Typ.PkgPath == pkg { // TODO: pkg is not full pkg path, only the package name
				if err := rg.Service.FillByAst(pkgAst); err != nil {
					log.Fatalf("fill service %s failed: %s", rg.Service.Typ.Indentity, err)
				}
			}
		}
	}
}
