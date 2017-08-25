package seeders

import (
	"github.com/goblog/database/factories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Seed() {
	db, err := gorm.Open("mysql", "root:password@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
	UserSeeder()
	ArticleSeeder()
	PostSeeder()

}
func UserSeeder() {
	for i := 1; i < 3; i++ {
		DB.Create(factory.UserFake())

	}

}
func ArticleSeeder() {
	for i := 1; i < 3; i++ {
		DB.Create(factory.ArticleFake())
	}

}
