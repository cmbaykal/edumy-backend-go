package studyEntity

type Study struct {
	ID       string `json:"id" bson:"_id"`
	UserId   string `json:"userId"`
	Lesson   string `json:"lesson"`
	CorrectA string `json:"correctA"`
	WrongA   string `json:"wrongA"`
	EmptyA   string `json:"emptyA"`
	Date     string `json:"date"`
}
