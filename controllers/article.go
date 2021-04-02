package controllers

import (
	"basic-api/config"
	"basic-api/forms"
	"basic-api/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FillArticle(db *gorm.DB, article *models.Article) error {
	if article.ID != 0 {
		err := db.Model(&models.User{}).Where("id = ?", article.UserID).Take(&article.User).Error
		if err != nil {
			return err
		}
		article.User.PrepareResponse()

		err = db.Model(&models.Comment{}).Where("article_id = ?", article.ID).Find(&article.Comments).Error
		if err != nil {
			return err
		}

		err = db.Model(&models.Opinion{}).Where("article_id = ?", article.ID).Find(&article.Opinion).Error
		return err
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
	err = db.Model(&models.Article{}).Where("id = ?", articleID).Take(&article).Error
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

	for _, article := range articles {
		FillArticle(db, &article)
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

	userID := r.Header.Get("user_id")
	var user models.User
	err2 := db.First(&user, &userID)
	if err2.Error != nil {
		log.Println("cannot get user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article := articleForm.PrepareArticle(user)
	err = db.Create(&article).Error
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

// PostArticles posts an article into the database
func EditArticle(w http.ResponseWriter, r *http.Request) {
	var articleForm forms.ArticleForm
	err := json.NewDecoder(r.Body).Decode(&articleForm)
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

	err = db.Model(&models.Article{}).Where("id = ?", article.ID).Updates(article).Error
	if err != nil {
		log.Println("cannot update user")
		w.WriteHeader(http.StatusForbidden)
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
