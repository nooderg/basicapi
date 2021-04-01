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
	r := mux.NewRouter()

	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/article", middlewares.Middleware(controllers.PostArticle)).Methods("POST")
	r.HandleFunc("/article/{id}", middlewares.Middleware(controllers.GetArticle)).Methods("GET")
	r.HandleFunc("/article/{id}/like", middlewares.Middleware(controllers.LikeArticle)).Methods("POST")
	r.HandleFunc("/article/{id}/dislike", middlewares.Middleware(controllers.DislikeArticle)).Methods("POST")

	r.HandleFunc("/article/{id}/comment", middlewares.Middleware(controllers.PostComment)).Methods("POST")
	r.HandleFunc("/article/{id}/comment/like", middlewares.Middleware(controllers.PostLikeComment)).Methods("POST")
	r.HandleFunc("/article/{id}/comment/dislike", middlewares.Middleware(controllers.PostDislikeComment)).Methods("POST")

	r.HandleFunc("/articles", middlewares.Middleware(controllers.ListArticles)).Methods("GET")

	r.HandleFunc("/login", middlewares.Middleware(controllers.HandleLogin)).Methods("POST")
	r.HandleFunc("/register", middlewares.Middleware(controllers.HandleRegister)).Methods("POST")
	r.HandleFunc("/profile", middlewares.Middleware(controllers.GetProfile)).Methods("GET")
	r.HandleFunc("/users/{id}", middlewares.Middleware(controllers.EditProfile)).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
