package main

import (
	"log"

	"github.com/podanypepa/cursor-pagination/db"
	"github.com/podanypepa/cursor-pagination/rest"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	}

	if err := rest.Server().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
