package models

type Email struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Content     string `json:"content"`
}
