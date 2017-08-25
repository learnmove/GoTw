package controllers

// github.com/goblog/controllers/user
import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/repositories"
)

type UserController struct {
	BaseController
	UserRepository    repositories.UserRepository
	ArticleRepository repositories.ArticleRepository
}

func (con UserController) init() {
	con.UserRepository = repositories.UserRepository{}
	con.ArticleRepository = repositories.ArticleRepository{}

}
func (con UserController) GetUser(c *gin.Context) {

	c.JSON(200, gin.H{
		"Message": con.UserRepository.GetUser(),
		"Article": con.ArticleRepository.GetArticle(),
	})
}
