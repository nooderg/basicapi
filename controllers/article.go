package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"basic-api/repository"
	"basic-api/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetArticle encodes an article
func GetArticle(w http.ResponseWriter, r *http.Request) {
	db := config.DBClient

	articleID := mux.Vars(r)["id"]
	article, err := repository.GetArticle(articleID)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	utils.FillArticle(db, article)

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetArticles encodes all the articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	db := config.DBClient

	articles, err := repository.GetArticles()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	for _, article := range articles {
		utils.FillArticle(db, &article)
	}

	err = json.NewEncoder(w).Encode(&articles)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PostArticle posts an article into the database
func PostArticle(w http.ResponseWriter, r *http.Request) {
	db := config.DBClient

	var articleForm forms.ArticleForm
	err := json.NewDecoder(r.Body).Decode(&articleForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("user_id")
	var user models.User
	err2 := db.First(&user, &userID)
	if err2.Error != nil {
		log.Println("cannot get user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := articleForm.PrepareArticle(user)
	err = repository.CreateArticle(&article)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// EditArticle posts an article into the database
func EditArticle(w http.ResponseWriter, r *http.Request) {
	db := config.DBClient

	var articleForm forms.ArticleForm
	err := json.NewDecoder(r.Body).Decode(&articleForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, _ := strconv.Atoi(r.Header.Get("user_id"))
	articleID := mux.Vars(r)["id"]
	var article models.Article
	err = db.Model(&models.Article{}).Where("id = ?", articleID).Take(&article).Error
	if err != nil {
		log.Println("article does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if uint(userID) != article.UserID {
		log.Println("you cannot edit this article!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	article.Title = articleForm.Title
	article.Content = articleForm.Content

	err = repository.UpdateArticle(&article)
	if err != nil {
		log.Println("cannot update user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	utils.FillArticle(db, &article)

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
