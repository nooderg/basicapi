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
	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var loginForm forms.LoginForm
	err = json.NewDecoder(r.Body).Decode(&loginForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user models.User
	res := db.Model(&models.User{}).Where("username = ?", loginForm.Username).Take(&user)
	if res.Error != nil {
		log.Println("user not found: " + loginForm.Username)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = user.CheckPassword(loginForm.Password)
	if err != nil {
		log.Println("wrong password for: " + loginForm.Username)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// waiting for jwt

	err = json.NewEncoder(w).Encode(true)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

	res := db.Model(&models.User{}).Where("username = ?", userForm.Username).Take(&models.User{})
	if res.Error == nil {
		log.Println("user already exists: " + userForm.Username)
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
