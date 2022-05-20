package main

import (
	"log"

	"github.com/podanypepa/cursor-pagination/rest"
)

func main() {
	if err := rest.Server().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
