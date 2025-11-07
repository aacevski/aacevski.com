package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {
	dir := "./static"

	templates = template.Must(template.ParseGlob("./static/templates/*.html"))

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/blog", BlogHandler).Methods("GET")
	router.HandleFunc("/books", BooksHandler).Methods("GET")
	router.HandleFunc("/rss", RSSHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:1337",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
