package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents the user in the DB
type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Dob      time.Time `json:"dob"`
	Email    string    `json:"email"`
	Username string    `json:"username" gorm:"column:username;not null;type:text"`
	Password string    `json:"password"  gorm:"column:password;not null;type:text"`
	Articles []Article
	Comments []Comment
	Opinions []Opinion
}

func (u *User) PrepareResponse() {
	u.Password = ""
}

func (u User) CheckPassword(pswrd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pswrd))
	return err
}
