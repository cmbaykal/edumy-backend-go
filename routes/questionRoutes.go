package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterQuestionRoutes(app *fiber.App) {
	class := app.Group("/question")
	class.Get("/add", controller.AddQuestion)
	class.Get("/delete", controller.DeleteQuestion)
	class.Get("/info", controller.QuestionInfo)
	class.Get("/classroomn", controller.ClassQuestions)
	class.Get("/user", controller.UserQuestions)
	class.Get("/feed", controller.QuestionsFeed)
	class.Get("/all", controller.AllQuestions)
}
