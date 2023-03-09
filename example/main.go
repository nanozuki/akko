package main

import "log"

func main() {
	srv, err := NewAkkoServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(srv.Run(":8080"))
}
