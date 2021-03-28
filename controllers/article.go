package controllers

import (
	"encoding/json"
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
	// encode all articles
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
