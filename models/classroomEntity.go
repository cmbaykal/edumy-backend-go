package models

type Classroom struct {
	ID        string   `json:"id,omitempty" bson:"_id,omitempty"`
	Lesson    string   `json:"lesson"`
	Name      string   `json:"name"`
	CreatorId string   `json:"creatorId" bson:"creatorId"`
	Users     []string `json:"users"`
}

type ClassroomResult struct {
	ID        string            `json:"id,omitempty" bson:"_id,omitempty"`
	Lesson    string            `json:"lesson"`
	Name      string            `json:"name"`
	CreatorId string            `json:"creatorId" bson:"creatorId"`
	Users     []User `json:"users"`
}

type MeetingClassroom struct {
	ID        string   `json:"id,omitempty" bson:"_id,omitempty"`
	Lesson    string   `json:"lesson"`
	Name      string   `json:"name"`
	CreatorId string   `json:"creatorId" bson:"creatorId"`
}
