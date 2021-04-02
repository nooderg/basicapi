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

// PostComment takes the CommentForm and fills the database
func PostComment(w http.ResponseWriter, r *http.Request) {
	var commentForm forms.CommentForm
	err := json.NewDecoder(r.Body).Decode(&commentForm)
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
	var user models.User
	err = db.Model(&models.User{}).Where("id = ?", userID).Take(&user).Error
	if err != nil {
		log.Println("user does not exist")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	articleID, _ := strconv.Atoi(mux.Vars(r)["id"])
	comment := commentForm.PrepareComment(uint(userID), uint(articleID))
	err = db.Create(&comment).Error
	if err != nil {
		log.Println("cannot create comment")
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

// EditComment takes the CommentForm and edits the matching comment
func EditComment(w http.ResponseWriter, r *http.Request) {
	var commentForm forms.CommentForm
	err := json.NewDecoder(r.Body).Decode(&commentForm)
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

	articleID, _ := strconv.Atoi(mux.Vars(r)["id"])
	commentID, _ := strconv.Atoi(mux.Vars(r)["commentID"])
	userID, _ := strconv.Atoi(r.Header.Get("user_id"))
	var comment models.Comment
	err = db.Model(&models.Comment{}).Where(&models.Comment{ID: uint(commentID), ArticleID: uint(articleID)}).Take(&comment).Error
	if err != nil {
		log.Println("cannot find comment on this article")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if uint(userID) != comment.UserID {
		log.Println("you cannot edit this comment!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	newComment := models.Comment{
		ID:        comment.ID,
		UserID:    comment.UserID,
		Title:     commentForm.Title,
		Content:   commentForm.Content,
		CreatedAt: comment.CreatedAt,
		ArticleID: comment.ArticleID,
	}

	err = db.Model(&models.Comment{}).Where("id = ?", newComment.ID).Updates(newComment).Error
	if err != nil {
		log.Println("cannot update user")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = json.NewEncoder(w).Encode(newComment)
	if err != nil {
		log.Println("cannot encode response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// DeleteComment removes a comment on
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentID, _ := strconv.Atoi(mux.Vars(r)["commentID"])

	db, err := config.InitDB()
	if err != nil {
		log.Println("cannot init db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := strconv.Atoi(r.Header.Get("user_id"))

	var comment models.Comment
	err = db.Model(&models.Comment{}).Where("id = ?", commentID).Take(&comment).Error
	if err != nil {
		log.Println("no opinion for this user on this article")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if comment.UserID != uint(userID) {
		log.Println("you cannot delete this comment!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = db.Model(&models.Comment{}).Where("id = ?", commentID).Delete(&comment).Error
	if err != nil {
		log.Println("no opinion for this user on this article")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(true)
}
