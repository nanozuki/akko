package akko

import (
	"fmt"
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
		fmt.Sprintf("Service: %s.%s", g.Service.Typ.PkgPath(), g.Service.Typ.Name()),
		fmt.Sprintf("\tDomain: %s", domain),
		fmt.Sprintf("\tBase: %s", g.Base),
		"\tMethods:",
	}
	for _, m := range g.Service.Methods {
		strs = append(strs, fmt.Sprintf("\t\t%s", m.Name()))
	}
	strs = append(strs, "\tProviders:")
	for p := range g.Providers {
		strs = append(strs, fmt.Sprintf("\t\t%s", p))
	}
	return strings.Join(strs, "\n")
}

type MountOption func(*RouteGroup)

type ServiceInfo struct {
	Typ     reflect.Type
	Methods []reflect.Type
}

func NewServiceInfo[T any](service T) *ServiceInfo {
	var svc T
	si := &ServiceInfo{
		Typ: reflect.TypeOf(svc),
	}
	si.Methods = make([]reflect.Type, 0, si.Typ.NumMethod())
	for i := 0; i < si.Typ.NumMethod(); i++ {
		si.Methods = append(si.Methods, si.Typ.Method(i).Type)
	}
	return si
}

type Provider[T any] func(*http.Request) (T, error)

type ProviderInfo struct {
	Func    reflect.Type
	Product reflect.Type
}

func NewProviderInfo[T any](p Provider[T]) ProviderInfo {
	var t T
	return ProviderInfo{
		Func:    reflect.TypeOf(p),
		Product: reflect.TypeOf(t),
	}
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
		var t T
		g.Providers[pi.Provide()] = &ProviderInfo{
			Func:    reflect.TypeOf(p),
			Product: reflect.TypeOf(t),
		}
	}
}

func PreRequest(fn func(*http.Request) (*http.Request, error)) MountOption {
	return func(*RouteGroup) {}
}

func PostResponse(fn func(*http.Response, error) (*http.Response, error)) MountOption {
	return func(*RouteGroup) {}
}

func WithMiddleware(middleware ...func(http.Handler) http.Handler) MountOption {
	return func(*RouteGroup) {}
}

var RouteGroups []*RouteGroup

func Mount[Srv any](baseRoute string, service Srv, options ...MountOption) {
	g := RouteGroup{
		Base:    baseRoute,
		Service: NewServiceInfo(service),
	}
	for _, opt := range options {
		opt(&g)
	}
	RouteGroups = append(RouteGroups, &g)
}
