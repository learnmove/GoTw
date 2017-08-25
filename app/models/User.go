package models

// . "github.com/goblog/app/models/user"

type User struct {
	BaseModel
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Articles []Article `json:"articles"`
}
