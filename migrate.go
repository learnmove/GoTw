package main

import (
	. "github.com/goblog/app/models/article"
	. "github.com/goblog/app/models/user"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:password@/goblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}

	defer db.Close()
	db.DropTableIfExists(&User{}, &Article{})
	db.AutoMigrate(&User{}, &Article{})

}
