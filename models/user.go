package models

// User represents the user in the database
type User struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// LoggedUser is used to return user data
type LoggedUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	City      string `json:"city"`
	Username  string `json:"username"`
}
