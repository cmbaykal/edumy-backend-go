package controllers

import (
	"github.com/cmbaykal/edumy-backend-go/database"
	"github.com/cmbaykal/edumy-backend-go/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserAuth(c *fiber.Ctx) error {
	return nil
}

func RegisterUser(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	userColl := db.Collection("userEntity")
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.SendString(err.Error())
	}

	result, err := userColl.InsertOne(c.Context(), user)

	if err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: result.InsertedID}}
	createdRecord := userColl.FindOne(c.Context(), filter)
	createdUser := &models.User{}
	createdRecord.Decode(createdUser)
	return c.JSON(createdUser)
}

func LoginUser(c *fiber.Ctx) error {
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	userColl := db.Collection("userEntity")
	updateCredentials := new(models.UpdateCredentials)

	if err := c.BodyParser(updateCredentials); err != nil {
		c.SendString(err.Error())
	}

	objectId, err := primitive.ObjectIDFromHex(updateCredentials.UserId)

	if err != nil {
		return c.SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: updateCredentials.Name},
				{Key: "bio", Value: updateCredentials.Bio},
			},
		},
	}

	err = userColl.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON("User update successful")
}

func ChangePasswordUser(c *fiber.Ctx) error {
	return nil
}

func UserInfo(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	userColl := db.Collection("userEntity")

	objectId, err := primitive.ObjectIDFromHex(c.Query("userId"))

	query := bson.D{{Key: "_id", Value: objectId}}

	if err != nil {
		return c.SendString(err.Error())
	}

	var user models.User
	err = userColl.FindOne(c.Context(), query).Decode(&user)

	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.GetInstance().Db
	userColl := db.Collection("userEntity")

	objectId, err := primitive.ObjectIDFromHex(c.Query("userId"))

	if err != nil {
		return c.SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: objectId}}

	result, err := userColl.DeleteOne(c.Context(), query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON("User Deleted Successfully")
}

func AllUsers(c *fiber.Ctx) error {
	query := bson.D{{}}
	collection := database.GetInstance().Db.Collection("userEntity")
	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	users := make([]models.User, 0)

	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users)
}
