package service

// go:generate go run github.com/nanozuki/akko/cmd/akko generate -o ../server

import (
	"github.com/nanozuki/akko"
)

type Service struct{}

func AkkoService() {
	akko.Mount("/v1", Service{}, akko.WithProvider(LoadUserByToken))
}
