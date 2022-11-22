package answerEntity

import (
	"github.com/cmbaykal/edumy-backend-go/data/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Answer struct {
	ID          string             `json:"id" bson:"_id"`
	QuestionId  string             `json:"questionId"`
	UserId      string             `json:"userId"`
	Description string             `json:"description"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Image       string             `json:"image"`
	Video       string             `json:"video"`
	UpVote      []string           `json:"upVote"`
	DownVote    []string           `json:"downVote"`
}

type AnswerInformation struct {
	ID          string             `json:"id" bson:"_id"`
	QuestionId  string             `json:"questionId"`
	User        userEntity.User    `json:"user"`
	Description string             `json:"description"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Image       string             `json:"image"`
	Video       string             `json:"video"`
	UpVote      []string           `json:"upVote"`
	DownVote    []string           `json:"downVote"`
}
