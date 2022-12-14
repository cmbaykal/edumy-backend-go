package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterAnswerRoutes(app *fiber.App) {
	class := app.Group("/answer")
	class.Get("/delete", controller.DeleteAnswer)
	class.Get("/upvote", controller.UpVoteAnswer)
	class.Get("/downvote", controller.DownVoteAnswer)
	class.Get("/question", controller.QuestionAnswers)
	class.Get("/user", controller.UserAnswers)
	class.Get("/classroom", controller.ClassAnswers)
	class.Get("/all", controller.AllAnswers)
}
