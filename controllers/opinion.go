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
)

// LikeArticle takes the LikeArticleForm
func RateArticle(w http.ResponseWriter, r *http.Request) {
	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var opinionForm forms.OpignionForm
	err = json.NewDecoder(r.Body).Decode(&opinionForm)
	if err != nil {
		log.Println("cannot decode request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	articleID, _ := strconv.Atoi(mux.Vars(r)["id"])
	userID, _ := strconv.Atoi(r.Header.Get("user_id"))

	var opinion models.Opinion
	err = db.Model(&models.Opinion{}).Where(&models.Opinion{UserID: uint(userID), ArticleID: uint(articleID)}).Take(&opinion).Error
	if err == nil {
		// Really GORM ???!!!
		err = db.Model(&models.Opinion{}).Where("id = ?", opinion.ID).Updates(map[string]interface{}{"like": opinionForm.Like}).Error
		if err != nil {
			log.Println("cannot update opinion")
			w.WriteHeader(http.StatusBadRequest)
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

		err = json.NewEncoder(w).Encode(article)
		if err != nil {
			log.Println("cannot encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		opinion := opinionForm.PrepareOpinion(uint(userID), uint(articleID))

		err = db.Create(&opinion).Error
		if err != nil {
			log.Println("cannot create opinion")
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

		err = json.NewEncoder(w).Encode(article)
		if err != nil {
			log.Println("cannot encode response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
