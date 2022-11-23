package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterMeetingRoutes(app *fiber.App) {
	class := app.Group("/meeting")
	class.Post("/schedule", controller.ScheduleMeeting)
	class.Post("/user", controller.UserMeetings)
}
