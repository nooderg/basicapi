package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"encoding/json"
	"log"
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
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var userForm forms.UserForm
	err = json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var userFound models.User
	res := db.Model(&userFound).Where("username = ?", userForm.Username).Find(&userFound)
	if res.RowsAffected != 0 {
		log.Println("user already exists")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := userForm.PrepareRegister()
	if err != nil {
		log.Println("cannot prepare register")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.Create(&user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetProfile takes the UserForm, edits the profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

// EditProfile takes the UserForm, edits the profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}
