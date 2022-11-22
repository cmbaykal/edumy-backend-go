package classroomRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cmbaykal/edumy-backend-go/controllers/classroom"
)

func Register(app *fiber.App) {
	class := app.Group("/class")
	class.Post("/add",classroomController.AddClass)
	class.Post("/assign",classroomController.AssignUser)
	class.Post("/leave",classroomController.LeaveClass)
	class.Post("/delete",classroomController.DeleteClass)
	class.Post("/info",classroomController.ClassInfo)
	class.Get("/all",classroomController.AllClasses)
}