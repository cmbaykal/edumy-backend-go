package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meeting struct {
	ID          string             `json:"id" bson:"_id"`
	ClassId     string             `json:"classId"`
	CreatorId   string             `json:"creatorId"`
	Description string             `json:"description"`
	Duration    int32              `json:"duration"`
	Date        primitive.DateTime `json:"date" bson:"date"`
}

type MeetingResult struct {
	ID          string             `json:"id" bson:"_id"`
	User        MeetingUser        `json:"user"`
	Classroom   MeetingClassroom   `json:"classroom"`
	Description string             `json:"description"`
	Duration    int32              `json:"duration"`
	Date        primitive.DateTime `json:"date" bson:"date"`
}
