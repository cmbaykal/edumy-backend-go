package meetingRoutes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/cmbaykal/edumy-backend-go/controllers/meeting"
)

func Register(app *fiber.App) {
	class := app.Group("/meeting")
	class.Get("/schedule",controller.ScheduleMeeting)
	class.Get("/user",controller.UserMeetings)
}