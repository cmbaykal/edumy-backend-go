package classroomEntity

import (
	"github.com/cmbaykal/edumy-backend-go/data/user"
)

type Classroom struct {
	ID        string   `json:"id" bson:"_id"`
	Lesson    string   `json:"lesson"`
	Name      string   `json:"name"`
	CreatorId string   `json:"creatorId"`
	Users     []string `json:"users"`
}

type ClassroomResult struct {
	ID        string            `json:"id" bson:"_id"`
	Lesson    string            `json:"lesson"`
	Name      string            `json:"name"`
	CreatorId string            `json:"creatorId"`
	Users     []userEntity.User `json:"users"`
}
