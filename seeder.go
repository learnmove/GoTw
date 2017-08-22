package main

import (
	"github.com/goblog/database/factories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db gorm.DB

func main() {

	db, err := gorm.Open("mysql", "root:ab789789@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()
	UserSeeder(db)
	FakerSeeder(db)

}
func UserSeeder(db *gorm.DB) {
	for i := 1; i < 3; i++ {
		db.Create(factory.UserFake())

	}
}
func FakerSeeder(db *gorm.DB) {
	for i := 1; i < 3; i++ {
		db.Create(factory.ArticleFake())
	}
}
