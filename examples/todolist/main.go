package main

//go:generate go run ../../cmd/akko/main.go -i ./service -o ./server -mod=github.com/nanozuki/akko/examples/todolist

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
