package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Question struct {
	ID          string             `json:"id" bson:"_id"`
	UserId      string             `json:"userId"`
	Lesson      string             `json:"lesson"`
	Description string             `json:"description"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Image       string             `json:"image"`
}

type QuestionInformation struct {
	ID          string             `json:"id" bson:"_id"`
	User        User    `json:"user"`
	Lesson      string             `json:"lesson"`
	Description string             `json:"description"`
	Date        primitive.DateTime `json:"date" bson:"date"`
	Image       string             `json:"image"`
}
