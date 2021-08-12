package main

import (
	// "fmt"
	"email_api/controllers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Post("/api/email", controllers.SendEmail)

	log.Fatal(app.Listen(":" + port))
}
