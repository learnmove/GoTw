package migrations

import (
	"github.com/goblog/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Migrate() {
	db, err := gorm.Open("mysql", "root:password@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
	DB.DropTableIfExists(&models.User{}, &models.Article{})
	DB.AutoMigrate(&models.User{}, &models.Article{})
	defer DB.Close()

}
