package repositories

import (
	"github.com/goblog/app/models"
)

type UserRepository struct{}

func (UserRepository UserRepository) GetUser() models.User {

	users := models.User{}
	// DB.Model(&users).Related(&articles, "Article")
	DB.Find(&users)
	return users
}
