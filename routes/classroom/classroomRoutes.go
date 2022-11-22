package classroomRoutes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/cmbaykal/edumy-backend-go/controllers/classroom"
)

func Register(app *fiber.App) {
	class := app.Group("/class")
	class.Get("/add",classroomController.AddClass)
	class.Get("/assign",classroomController.AssignUser)
	class.Get("/leave",classroomController.LeaveClass)
	class.Get("/delete",classroomController.DeleteClass)
	class.Get("/info",classroomController.ClassInfo)
	class.Get("/all",classroomController.AllClasses)
}