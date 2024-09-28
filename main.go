package main

import (
	"log"

	"github.com/colt005/wiki/server"
)

func main() {

	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	s.RegisterRoutes()

	s.Start()
}
