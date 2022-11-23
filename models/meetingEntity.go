package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meeting struct {
	ID          string             `json:"id" bson:"_id"`
	ClassId     string             `json:"classId"`
	CreatorId   string             `json:"creatorId"`
	Description string             `json:"description"`
	Duration    string             `json:"duration"`
	Date        primitive.DateTime `json:"date" bson:"date"`
}

type MeetingResult struct {
	ID          string                    `json:"id" bson:"_id"`
	User        User           `json:"user"`
	Classroom   Classroom `json:"classroom"`
	Description string                    `json:"description"`
	Duration    string                    `json:"duration"`
	Date        primitive.DateTime        `json:"date" bson:"date"`
}
