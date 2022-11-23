package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthToken struct {
	AuthToken  string             `json:"token"`
	ExpireTime primitive.DateTime `json:"expireTime" bson:"expireTime"`
}
