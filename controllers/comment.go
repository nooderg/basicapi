package controllers

import (
	"basic-api/config"
	"basic-api/models"
	"encoding/json"
	"log"
	"net/http"
)

// PostComment takes the CommentForm and fills the database
func PostComment(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)

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

	res := db.Model(&comment).Where("id = ?", comment.ID).Take(&comment)
	if res.Error != nil {
		log.Println("comment does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Println(comment)

	err = json.NewEncoder(w).Encode(&comment)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PostLikeComment takes the CommentForm and fills the database
func PostLikeComment(w http.ResponseWriter, r *http.Request) {
	// do stuff
}

// PostDislikeComment takes the CommentForm and fills the database
func PostDislikeComment(w http.ResponseWriter, r *http.Request) {
	// do stuff
}
