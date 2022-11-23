package routes

import (
	controller "github.com/cmbaykal/edumy-backend-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	class := app.Group("/user")
	class.Get("/auth/token", controller.UserAuth)
	class.Get("/register", controller.RegisterUser)
	class.Get("/login", controller.LoginUser)
	class.Get("/update", controller.UpdateUser)
	class.Get("/changePassword", controller.ChangePasswordUser)
	class.Get("/info", controller.UserInfo)
	class.Get("/delete", controller.DeleteUser)
	class.Get("/all", controller.AllUsers)
}
