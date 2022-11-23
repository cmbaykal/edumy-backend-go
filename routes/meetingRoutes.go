package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterMeetingRoutes(app *fiber.App) {
	class := app.Group("/meeting")
	class.Get("/schedule", controller.ScheduleMeeting)
	class.Get("/user", controller.UserMeetings)
}
