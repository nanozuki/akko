package akko

import "net/http"

type MountOptions struct{}

type MountOption func(*MountOptions)

func WithProvider[T any](fn func(*http.Request) (T, error)) MountOption {
	return func(*MountOptions) {}
}

func PreRequest(fn func(*http.Request) (*http.Request, error)) MountOption {
	return func(*MountOptions) {}
}

func PostResponse(fn func(*http.Response, error) (*http.Response, error)) MountOption {
	return func(*MountOptions) {}
}

func WithMiddleware(middleware ...func(http.Handler) http.Handler) MountOption {
	return func(*MountOptions) {}
}

func Mount(route string, service any, options ...MountOption) {}
