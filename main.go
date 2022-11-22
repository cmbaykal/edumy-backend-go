package main

import (
	"log"
	
	"github.com/cmbaykal/edumy-backend-go/database"
	"github.com/cmbaykal/edumy-backend-go/routes/answer"
	"github.com/cmbaykal/edumy-backend-go/routes/classroom"
	"github.com/cmbaykal/edumy-backend-go/routes/meeting"
	"github.com/cmbaykal/edumy-backend-go/routes/question"
	"github.com/cmbaykal/edumy-backend-go/routes/study"
	"github.com/cmbaykal/edumy-backend-go/routes/user"
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

func registerRoutes(app *fiber.App){
	answerRoutes.Register(app)
	classroomRoutes.Register(app)
	meetingRoutes.Register(app)
	questionRoutes.Register(app)
	studyRoutes.Register(app)
	userRoutes.Register(app)
}
