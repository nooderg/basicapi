package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"encoding/json"
	"net/http"
)

// HandleLogin checks the email and password and returns the token
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

// HandleRegister takes the UserForm, registers the user
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	var userForm forms.UserForm
	err = json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		panic(err)
	}

	//
	// Make a method in UserForm (PrepareRegister) to check potential errors in the form + hash password + return a User
	//

	user := models.User{
		FirstName: userForm.FirstName,
		LastName:  userForm.LastName,
		City:      userForm.City,
		Dob:       userForm.Dob.Format("2006-01-02"),
		Username:  userForm.Username,
		Password:  userForm.Password,
	}

	db.Create(&user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		panic(err)
	}
}

// GetProfile takes the UserForm, edits the profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// log.Println(user)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		panic(err)
	}

}

// EditProfile takes the UserForm, edits the profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}
