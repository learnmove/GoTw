package repositories

import (
	"github.com/goblog/app/models"
)

type ArticleRepository struct{}

func (ArticleRepository ArticleRepository) GetArticle() models.Article {

	articles := models.Article{}
	// DB.Model(&users).Related(&articles, "Article")
	DB.Find(&articles)
	return articles
}
