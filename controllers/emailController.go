package controllers

import (
	m "email_api/models"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SendEmail(c *fiber.Ctx) error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := new(m.Email)

	erro := c.BodyParser(e)

	if erro != nil {
		log.Fatal(erro)
		return c.SendString("Email was not sent.")
	}

	if e.Email == "" {
		return c.SendString("Must enter an email to send to.")
	}

	if e.Content == "" {
		return c.SendString("Must enter some content.")
	}

	// Sender info
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	to := []string{
		"millurr0@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Initiate message to send with User email, User phone number, and User content
	message := []byte(e.Email + " | " + e.PhoneNumber + " | " + e.Content)

	// Authenticate
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	er := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if er != nil {
		fmt.Println(er)
		return c.SendString("Email was not able to send.")
	}

	return c.SendString("Email sent successfully.")
}
