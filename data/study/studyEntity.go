package studyEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Study struct {
	ID       string             `json:"id" bson:"_id"`
	UserId   string             `json:"userId"`
	Lesson   string             `json:"lesson"`
	CorrectA string             `json:"correctA"`
	WrongA   string             `json:"wrongA"`
	EmptyA   string             `json:"emptyA"`
	Date     primitive.DateTime `json:"date" bson:"date"`
}
