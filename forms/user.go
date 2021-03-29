package forms

import (
	"basic-api/models"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func (u UserForm) PrepareRegister() (*models.User, error) {
	hashedPassword, err := (bcrypt.GenerateFromPassword([]byte(u.Password), 10))
	if err != nil {
		return nil, err
	}
	user := models.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		City:      u.City,
		Dob:       u.Dob.Format("2006-01-02"),
		Username:  u.Username,
		Password:  string(hashedPassword),
	}
	return &user, nil
}
