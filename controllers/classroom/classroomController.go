package classroomController

import (
	"fmt"

	classroomEntity "github.com/cmbaykal/edumy-backend-go/data/classroom"
	userEntity "github.com/cmbaykal/edumy-backend-go/data/user"
	"github.com/cmbaykal/edumy-backend-go/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddClass(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	classColl := db.Collection("classroom")

	class := new(classroomEntity.Classroom)

	if err := c.BodyParser(class); err != nil {
		return c.SendString(err.Error())
	}

	class.Users = append(class.Users, class.CreatorId)

	result, err := classColl.InsertOne(c.Context(), class)

	if err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := classColl.FindOne(c.Context(), filter)
	createdClass := &classroomEntity.Classroom{}
	createdRecord.Decode(createdClass)
	return c.JSON(createdClass)
}

func AssignUser(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	classColl := db.Collection("classroom")
	userColl := db.Collection("userEntity")

	queryClassId := c.Query("classId")
	queryUserMail := c.Query("userMail")

	classQuery := bson.D{{Key: "_id", Value: queryClassId}}
	userQuery := bson.D{{Key: "mail", Value: queryUserMail}}

	var class classroomEntity.Classroom
	var user userEntity.User

	err := classColl.FindOne(c.Context(), classQuery).Decode(&class)

	if err != nil {
		return c.SendStatus(500)
	}

	err = userColl.FindOne(c.Context(), userQuery).Decode(&user)

	if err != nil {
		return c.SendStatus(500)
	}

	if user.Role == userEntity.Student {
		class.Users = append(class.Users, user.ID)

		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "users", Value: class.Users},
				},
			},
		}

		err = classColl.FindOneAndUpdate(c.Context(), classQuery, update).Err()

		if err != nil {
			return c.SendStatus(500)
		}

		return c.JSON(class)
	}

	return c.Status(404).JSON("User is not a student")
}

func LeaveClass(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	classColl := db.Collection("classroom")
	userColl := db.Collection("userEntity")

	queryClassId := c.Query("classId")
	queryUserMail := c.Query("userMail")

	classQuery := bson.D{{Key: "_id", Value: queryClassId}}
	userQuery := bson.D{{Key: "mail", Value: queryUserMail}}

	var class classroomEntity.Classroom
	var user userEntity.User

	err := classColl.FindOne(c.Context(), classQuery).Decode(&class)

	if err != nil {
		return c.SendStatus(500)
	}

	err = userColl.FindOne(c.Context(), userQuery).Decode(&user)

	if err != nil {
		return c.SendStatus(500)
	}

	var foundIndex int

	for index, item := range class.Users {
		if item == user.ID {
			foundIndex = index
			break
		}
	}

	class.Users = append(class.Users[:foundIndex], class.Users[foundIndex+1:]...)

	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "users", Value: class.Users},
			},
		},
	}

	err = classColl.FindOneAndUpdate(c.Context(), classQuery, update).Err()

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(class)
}

func DeleteClass(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	classColl := db.Collection("classroom")

	queryParam := c.Query("classId")
	query := bson.D{{Key: "_id", Value: queryParam}}

	result, err := classColl.DeleteOne(c.Context(), query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("Classroom Deleted Successfully")
}

func ClassInfo(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	classColl := db.Collection("classroom")
	userColl := db.Collection("userEntity")

	objectId, err := primitive.ObjectIDFromHex(c.Query("classId"))

	if err != nil {
		return c.SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: objectId}}

	var class classroomEntity.Classroom
	err = classColl.FindOne(c.Context(), query).Decode(&class)

	if err != nil {
		return c.SendString(err.Error())
	}

	var users []userEntity.User

	for _, item := range class.Users {
		query := bson.D{{Key: "_id", Value: item}}

		var user userEntity.User
		err := userColl.FindOne(c.Context(), query).Decode(&user)
		if err != nil {
			fmt.Println("Second")
			return c.SendString(err.Error())
		}
		users = append(users, user)
	}

	result := classroomEntity.ClassroomResult{
		ID:        class.ID,
		Lesson:    class.Lesson,
		Name:      class.Name,
		CreatorId: class.CreatorId,
		Users:     users,
	}

	return c.JSON(result)
}

func AllClasses(c *fiber.Ctx) error {
	query := bson.D{{}}
	collection := database.GetInstance().Db.Collection("classroom")
	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	classrooms := make([]classroomEntity.Classroom, 0)

	if err := cursor.All(c.Context(), &classrooms); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(classrooms)
}
