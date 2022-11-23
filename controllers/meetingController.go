package controllers

import (
	"github.com/cmbaykal/edumy-backend-go/database"
	"github.com/cmbaykal/edumy-backend-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ScheduleMeeting(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	meetingColl := db.Collection("meeting")
	meeting := new(models.Meeting)

	if err := c.BodyParser(meeting); err != nil {
		return c.SendString(err.Error())
	}

	result, err := meetingColl.InsertOne(c.Context(), meeting)

	if err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := meetingColl.FindOne(c.Context(), filter)
	createdMeeting := &models.Meeting{}
	createdRecord.Decode(createdMeeting)
	return c.JSON(createdMeeting)
}

func UserMeetings(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	meetingColl := db.Collection("meeting")
	userColl := db.Collection("userEntity")
	classColl := db.Collection("classroom")

	query := bson.D{{Key: "_id", Value: c.Query("userId")}}

	var user models.User
	err := userColl.FindOne(c.Context(), query).Decode(&user)

	if err != nil {
		return c.SendString(err.Error())
	}

	meetingQuery := bson.D{{Key: "creatorId", Value: c.Query("userId")}}

	cursor, err := meetingColl.Find(c.Context(), meetingQuery)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	meetings := make([]models.Meeting, 0)

	if err := cursor.All(c.Context(), &meetings); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	result := make([]models.MeetingResult, 0)

	for _, item := range meetings {
		classId, err := primitive.ObjectIDFromHex(item.ClassId)

		if err != nil {
			return c.SendString(err.Error())
		}

		creatorQuery := bson.D{{Key: "_id", Value: item.CreatorId}}
		classQuery := bson.D{{Key: "_id", Value: classId}}

		var meetingUser models.MeetingUser
		err = userColl.FindOne(c.Context(), creatorQuery).Decode(&meetingUser)

		if err != nil {
			return c.SendString(err.Error())
		}

		var class models.MeetingClassroom
		err = classColl.FindOne(c.Context(), classQuery).Decode(&class)

		if err != nil {
			return c.SendString(err.Error())
		}

		resultItem := models.MeetingResult{
			ID:          item.ID,
			User:        meetingUser,
			Classroom:   class,
			Description: item.Description,
			Duration:    item.Duration,
			Date:        item.Date,
		}

		result = append(result, resultItem)
	}

	return c.JSON(result)
}
