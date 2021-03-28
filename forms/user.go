package forms

import (
	"basic-api/models"
	"time"
)

// UserForm is used to edit or create an user
type UserForm struct {
	FirstName       string    `json:"firstname"`
	LastName        string    `json:"lastname"`
	Dob             time.Time `json:"dob"`
	City            string    `json:"city"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirm_password"`
}

// LoginForm is used to log the user in
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*UserForm) PrepareRegister() (*models.User, error) {
	return nil, nil
}
