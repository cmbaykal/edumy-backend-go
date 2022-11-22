package questionEntity

import (
	"github.com/cmbaykal/edumy-backend-go/data/user"
)

type Question struct {
	ID          string `json:"id" bson:"_id"`
	UserId      string `json:"userId"`
	Lesson      string `json:"lesson"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Image       string `json:"image"`
}

type QuestionInformation struct {
	ID          string          `json:"id" bson:"_id"`
	User        userEntity.User `json:"user"`
	Lesson      string          `json:"lesson"`
	Description string          `json:"description"`
	Date        string          `json:"date"`
	Image       string          `json:"image"`
}
