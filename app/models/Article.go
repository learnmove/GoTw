package models

// . "github.com/goblog/app/models/article"

type Article struct {
	BaseModel
	Title   string
	Content string
	UserID  int
}
