package main

import (
	"log"

	"github.com/itpkg/chaos/web"
)

func main() {
	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
