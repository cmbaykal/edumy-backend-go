package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	class := app.Group("/study")
	class.Get("/add", controller.AddStudy)
	class.Get("/delete", controller.DeleteStudy)
	class.Get("/user", controller.UserStudies)
	class.Get("/classroom", controller.ClassStudies)
	class.Get("/feed", controller.StudiesFeed)
	class.Get("/all", controller.AllStudies)
}
