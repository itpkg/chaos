package main

import (
	"log"

	_ "github.com/itpkg/chaos/engines/cms"
	_ "github.com/itpkg/chaos/engines/hr"
	_ "github.com/itpkg/chaos/engines/ops/mail"
	_ "github.com/itpkg/chaos/engines/ops/vpn"
	_ "github.com/itpkg/chaos/engines/reading"
	_ "github.com/itpkg/chaos/engines/team"
	"github.com/itpkg/chaos/web"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
