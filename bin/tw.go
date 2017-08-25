package main

import (
	// . "github.com/goblog/app/models"
	"flag"
	"log"
	// "github.com/jinzhu/gorm"
	"github.com/goblog/database/migrations"
	"github.com/goblog/database/seeders"
	"github.com/goblog/stubs"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	do := flag.String("do", "", "You Want Do")
	name := flag.String("name", "", "You File Name")

	flag.Parse()

	switch *do {
	case "migrate":
		migrate()

	case "seed":
		seeder()
	case "make:model":
		if name != nil {
			stubModel(*name)

		}
	case "make:controller":
		if name != nil {
			stubController(*name)

		}
	case "make:repository":
		if name != nil {
			stubRepository(*name)

		}
	case "migrate:seed":
		migrate()
		seeder()

	default:
		log.Fatal("please input like 'go run tw.go -do migrate/seed/migrate:seed/make:controller/make:model -name Article'   ")
	}

}

func migrate() {
	migrations.Migrate()
}

func seeder() {
	seeders.Seed()
}
func stubModel(name string) {
	stubs.CreateModel(name)
}
func stubController(name string) {
	stubs.CreateController(name)
}
func stubRepository(name string) {
	stubs.CreateRepository(name)

}
