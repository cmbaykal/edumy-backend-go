package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      string             `json:"id" bson:"_id"`
	Role    UserRole           `json:"role"`
	Mail    string             `json:"mail"`
	Birth   primitive.DateTime `json:"birth" bson:"birth"`
	Name    string             `json:"name"`
	Bio     string             `json:"bio"`
	Classes []string           `json:"classes"`
}

type MeetingUser struct {
	ID      string             `json:"id" bson:"_id"`
	Mail    string             `json:"mail"`
	Name    string             `json:"name"`
	Bio     string             `json:"bio"`
}

type UserEntity struct {
	User
	Pass string `json:"pass"`
}

type UserCredentials struct {
	Mail string `json:"mail"`
	Pass string `json:"pass"`
}

type UpdateCredentials struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
}

type PasswordCredentials struct {
	UserId  string `json:"userId"`
	OldPass string `json:"oldPass"`
	NewPass string `json:"newPass"`
}

type UserRole string

const (
	Student UserRole = "Student"
	Teacher          = "Teacher"
)
