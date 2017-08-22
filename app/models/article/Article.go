package Article

// . "github.com/goblog/app/models/article"

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
	UserID  int
}
