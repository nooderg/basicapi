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

	r.HandleFunc("/article", middlewares.JWTVerify(controllers.PostArticle)).Methods("POST")
	r.HandleFunc("/article/{id}", controllers.GetArticle).Methods("GET")
	r.HandleFunc("/article/{id}", middlewares.JWTVerify(controllers.EditArticle)).Methods("PUT")
	r.HandleFunc("/article/{id}/opinion", middlewares.JWTVerify(controllers.RateArticle)).Methods("POST")
	r.HandleFunc("/article/{id}/opinion", middlewares.JWTVerify(controllers.DeleteOpinion)).Methods("DELETE")

	r.HandleFunc("/articles", controllers.ListArticles).Methods("GET")

	r.HandleFunc("/article/{id}/comment", middlewares.JWTVerify(controllers.PostComment)).Methods("POST")
	r.HandleFunc("/article/{id}/comment/{commentID}", middlewares.JWTVerify(controllers.EditComment)).Methods("PUT")
	r.HandleFunc("/article/{id}/comment/{commentID}", middlewares.JWTVerify(controllers.DeleteComment)).Methods("DELETE")

	r.HandleFunc("/login", controllers.HandleLogin).Methods("POST")
	r.HandleFunc("/register", controllers.HandleRegister).Methods("POST")
	r.HandleFunc("/me", middlewares.JWTVerify(controllers.GetMe)).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/users/{id}", middlewares.JWTVerify(controllers.EditProfile)).Methods("PUT")

	log.Println("Server running on localhost:8080!")

	log.Fatal(http.ListenAndServe(":8080", r))
}
