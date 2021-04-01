package models

import "time"

// User represents the user in the DB
type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Dob      time.Time `json:"dob"`
	Username string    `json:"username" gorm:"column:username;not null;type:text"`
	Password string    `json:"password"  gorm:"column:password;not null;type:text"`
}

// LoggedUser is used to return user data
type LoggedUser struct {
	ID       uint      `json:"id"`
	Dob      time.Time `json:"dob"`
	Username string    `json:"username"`
}
