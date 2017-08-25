package repositories

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {

	db, err := gorm.Open("mysql", "root:password@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	DB = db
}
