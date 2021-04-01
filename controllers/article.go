package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Fills the article's author
func FillArticle(db *gorm.DB, article *models.Article) error {
	if article.ID != 0 {
		return db.Model(&models.User{}).Where("id = ?", article.UserID).Take(&article.Author).Error
	}
	return nil
}

// GetArticle enocodes an article
func GetArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]

	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var article models.Article
	err = db.Model(&models.Article{}).Where("id = ?", articleID).Find(&article).Error
	if err != nil {
		log.Println("article does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	FillArticle(db, &article)

	err = json.NewEncoder(w).Encode(&article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetArticles enocodes all the articles
func ListArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article

	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = db.Model(&models.Article{}).Find(&articles).Error
	if err != nil {
		log.Println("article does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(&articles)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// PostArticles posts an article into the database
func PostArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var articleForm forms.ArticleForm
	err = json.NewDecoder(r.Body).Decode(&articleForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := articleForm.PrepareArticle()

	err = db.Table("articles").Create(&article).Error
	if err != nil {
		log.Println("cannot create article")
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// LikeArticle takes the LikeArticleForm
func LikeArticle(w http.ResponseWriter, r *http.Request) {
	// fill database
}

// DislikeArticle takes the LikeArticleForm
func DislikeArticle(w http.ResponseWriter, r *http.Request) {
	// fill database
}
