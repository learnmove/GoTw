package factory

import (
	. "github.com/goblog/app/models"
	"github.com/icrowley/fake"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func UserFake() interface{} {

	user := User{
		Name: fake.FirstName(),
		Age:  rand.Intn(100),
	}
	return &user

}
func ArticleFake() interface{} {
	article := Article{
		Title:   fake.Title(),
		Content: fake.Paragraph(),
		UserID:  rand.Intn(10) + 1,
	}

	return &article
}
