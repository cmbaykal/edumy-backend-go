package classroomController

import (
	"github.com/cmbaykal/edumy-backend-go/data/classroom"
	"github.com/cmbaykal/edumy-backend-go/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func AddClass(c *fiber.Ctx) error {
	return nil
}

func AssignUser(c *fiber.Ctx) error {
	return nil
}

func LeaveClass(c *fiber.Ctx) error {
	return nil
}

func DeleteClass(c *fiber.Ctx) error {
	return nil
}

func ClassInfo(c *fiber.Ctx) error {
	return nil
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
