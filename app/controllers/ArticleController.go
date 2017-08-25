package controllers

// github.com/goblog/controllers/user
import (
	"github.com/gin-gonic/gin"
	"github.com/goblog/app/repositories"
)

type ArticleController struct {
	BaseController
	UserRepository    repositories.UserRepository
	ArticleRepository repositories.ArticleRepository
}

func (con ArticleController) init() {
	con.UserRepository = repositories.UserRepository{}
	con.ArticleRepository = repositories.ArticleRepository{}
}
func (con ArticleController) GetArticle(c *gin.Context) {

	c.JSON(200, gin.H{
		"uSER":    con.UserRepository.GetUser(),
		"ARITLCE": con.ArticleRepository.GetArticle(),
	})
}
