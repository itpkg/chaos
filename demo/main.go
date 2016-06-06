package main

import (
	"log"

	_ "github.com/itpkg/chaos/engines/auth"
	_ "github.com/itpkg/chaos/engines/cms"
	_ "github.com/itpkg/chaos/engines/hr"
	_ "github.com/itpkg/chaos/engines/ops"
	_ "github.com/itpkg/chaos/engines/reading"
	"github.com/itpkg/chaos/web"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	if err := web.Run(); err != nil {
		log.Fatal(err)
	}
}
