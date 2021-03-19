package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonoy30/practice-rest/handlers"
	"github.com/tonoy30/practice-rest/middlewares"
)

const (
	PORT = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.Headers("x-content-type-options", "nosniff")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/", middlewares.Logging(handlers.Home)).Methods(http.MethodGet)
	r.HandleFunc("/contact", middlewares.Logging(handlers.Contact)).Methods(http.MethodGet, http.MethodPost)
	articleRoute(r)
	log.Fatal(http.ListenAndServe(PORT, r))
}
func articleRoute(r *mux.Router) {
	r.HandleFunc("/article", middlewares.Logging(handlers.AllArticle)).Methods(http.MethodGet)
	r.HandleFunc("/article", middlewares.Logging(handlers.CreateArticle)).Methods(http.MethodPost)
	r.HandleFunc("/article/{id}", middlewares.Logging(handlers.GetArticleById)).Methods(http.MethodGet)
	r.HandleFunc("/article/{id}", middlewares.Logging(handlers.UpdateArticleById)).Methods(http.MethodPut)
	r.HandleFunc("/article/{id}", middlewares.Logging(handlers.DeleteArticleById)).Methods(http.MethodDelete)
}
