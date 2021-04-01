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
	err = db.Model(&models.User{}).Where("username = ?", loginForm.Username).Take(&user).Error
	if err != nil {
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

	err = db.Model(&models.User{}).Where("username = ?", userForm.Username).Take(&models.User{}).Error
	if err == nil {
		log.Println("user already exists: " + userForm.Username)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := userForm.PrepareUser()
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

	user.PrepareResponse()

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetProfile takes the UserForm, edits the profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get user ID with jwt
	//Temporary solution
	id := 1
	var user models.User
	err = db.Model(&models.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		log.Println("user does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user.PrepareResponse()
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// EditProfile takes the UserForm, edits the profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var userForm forms.UserForm
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userForm.Password != userForm.ConfirmPassword {
		log.Println("passwords are not matching")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Get user ID with jwt AYMERIC IM LOOKING AT YOU
	//Temporary solution
	id := 1
	var user models.User
	err = db.Model(&models.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		log.Println("user does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	newUser, err := userForm.PrepareUser()
	newUser.ID = uint(id)

	newUser.PrepareResponse()

	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
