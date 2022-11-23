package main

import (
	"log"

	"github.com/cmbaykal/edumy-backend-go/database"
	routes "github.com/cmbaykal/edumy-backend-go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	registerRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func registerRoutes(app *fiber.App) {
	routes.RegisterAnswerRoutes(app)
	routes.RegisterClassroomRoutes(app)
	routes.RegisterMeetingRoutes(app)
	routes.RegisterMeetingRoutes(app)
	routes.RegisterQuestionRoutes(app)
	routes.RegisterUserRoutes(app)
}
