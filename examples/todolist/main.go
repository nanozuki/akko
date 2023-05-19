package main

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
