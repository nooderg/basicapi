package controllers

import (
	"basic-api/config"
	"basic-api/models"
	"encoding/json"
	"log"
	"net/http"
)

// GetArticle enocodes an article
func GetArticle(w http.ResponseWriter, r *http.Request) {
	//Dumb data
	article := struct {
		ID   string
		text string
	}{
		"333",
		"Texttttt",
	}

	//How you send stuff in the response
	json.NewEncoder(w).Encode(article)
}

// GetArticles enocodes all the articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)

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

	res := db.Model(&article).Where("id = ?", article.ID).Take(&article)
	if res.Error != nil {
		log.Println("article does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	log.Println(article)

	err = json.NewEncoder(w).Encode(&article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PostArticles posts an array of articles into the database
func PostArticles(w http.ResponseWriter, r *http.Request) {
	// fill database
}

// LikeArticle takes the LikeArticleForm
func LikeArticle(w http.ResponseWriter, r *http.Request) {
	// fill database
}

// DislikeArticle takes the LikeArticleForm
func DislikeArticle(w http.ResponseWriter, r *http.Request) {
	// fill database
}
