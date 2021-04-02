package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PostComment takes the CommentForm and fills the database
func PostComment(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var commentForm forms.CommentForm
	err = json.NewDecoder(r.Body).Decode(&commentForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")
	var user models.User
	err = db.Model(&models.User{}).Where("id = ?", userID).Take(&user).Error
	if err != nil {
		log.Println("user does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	articleID := mux.Vars(r)["id"]
	var article models.Article
	err = db.Model(&models.Article{}).Where("id = ?", articleID).Take(&article).Error
	if err != nil {
		log.Println("article does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	comment := commentForm.PrepareComment(user, article)
	article.Comments = append(article.Comments, comment)

	err = db.Create(&comment).Error
	if err != nil {
		log.Println("cannot create comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	FillArticle(db, &article)

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
