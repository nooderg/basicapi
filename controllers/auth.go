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

	res := db.Model(&models.User{}).Where("username = ?", userForm.Username).Find(&models.User{})
	if res.RowsAffected != 0 {
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

	err = db.Create(&user).Error
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetProfile takes the UserForm, edits the profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := db.Model(&user).Where("id = ?", user.ID).Take(&user)
	if res.Error != nil {
		log.Println("user does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Println(user)

	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// EditProfile takes the UserForm, edits the profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	// do stuff
}
