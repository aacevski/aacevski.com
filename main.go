package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"aacevski.com/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var templates *template.Template

func main() {
	godotenv.Load()

	dir := "./static"

	templates = template.Must(template.New("").Funcs(template.FuncMap{
		"divf": func(a, b int) float64 {
			return float64(a) / float64(b)
		},
	}).ParseGlob("./static/templates/*.html"))

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(w, r, templates)
	}).Methods("GET")
	router.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		handlers.Blog(w, r, templates)
	}).Methods("GET")
	router.HandleFunc("/blog/{slug}", func(w http.ResponseWriter, r *http.Request) {
		handlers.BlogPostPage(w, r, templates)
	}).Methods("GET")
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.Books(w, r, templates)
	}).Methods("GET")
	router.HandleFunc("/rss.xml", func(w http.ResponseWriter, r *http.Request) {
		handlers.RSS(w, r)
	}).Methods("GET")
	// Keep /rss for backwards compatibility
	router.HandleFunc("/rss", func(w http.ResponseWriter, r *http.Request) {
		handlers.RSS(w, r)
	}).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:1337",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
