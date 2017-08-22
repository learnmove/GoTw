package User

// . "github.com/goblog/app/models/user"

import (
	. "github.com/goblog/app/models/article"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Articles []Article
}
