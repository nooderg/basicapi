package main

import (
	"basic-api/config"
	"basic-api/controllers"
	"basic-api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server...")
	r := mux.NewRouter()

	log.Println("Connecting to database...")
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database!")

	r.HandleFunc("/article", controllers.PostArticle).Methods("POST")
	r.HandleFunc("/article/{id}", controllers.GetArticle).Methods("GET")
	r.HandleFunc("/article/{id}/like", controllers.LikeArticle).Methods("POST")
	r.HandleFunc("/article/{id}/dislike", controllers.DislikeArticle).Methods("POST")

	r.HandleFunc("/article/{id}/comment", controllers.PostComment).Methods("POST")
	r.HandleFunc("/article/{id}/comment/like", controllers.PostLikeComment).Methods("POST")
	r.HandleFunc("/article/{id}/comment/dislike", controllers.PostDislikeComment).Methods("POST")

	r.HandleFunc("/articles", controllers.ListArticles).Methods("GET")

	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	r.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
	r.HandleFunc("/me", middlewares.JWTVerify(controllers.GetProfile)).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/users/{id}", middlewares.JWTVerify(controllers.EditProfile)).Methods("PUT")

	log.Println("Server running!")

	log.Fatal(http.ListenAndServe(":8080", r))
}
