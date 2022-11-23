package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterClassroomRoutes(app *fiber.App) {
	class := app.Group("/class")
	class.Post("/add", controller.AddClass)
	class.Post("/assign", controller.AssignUser)
	class.Post("/leave", controller.LeaveClass)
	class.Post("/delete", controller.DeleteClass)
	class.Post("/info", controller.ClassInfo)
	class.Get("/all", controller.AllClasses)
}
