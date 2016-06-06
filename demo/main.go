package main

import (
	"log"

	_ "github.com/itpkg/chaos/engines/auth"
	_ "github.com/itpkg/chaos/engines/cms"
	_ "github.com/itpkg/chaos/engines/hr"
	_ "github.com/itpkg/chaos/engines/ops"
	_ "github.com/itpkg/chaos/engines/reading"
	"github.com/itpkg/chaos/web"
)

func main() {
	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
