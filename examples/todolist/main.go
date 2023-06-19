package main

//go:generate go run ../../cmd/akko/main.go -i ./service -o ./server -s Service -mod github.com/nanozuki/akko/examples/todolist/service

import (
	"log"

	"github.com/nanozuki/akko/examples/todolist/server"
	"github.com/nanozuki/akko/examples/todolist/service"
)

func main() {
	service := service.Service{}
	server := server.NewServer(&service)
	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal(err)
	}
}
