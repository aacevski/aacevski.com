package handlers

import (
	"html/template"
	"log"
	"net/http"

	"aacevski.com/pkg/koreader"
	"github.com/joho/godotenv"
)

func Books(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	godotenv.Load()

	log.Println("Fetching book statistics from KOReader...")
	readingStats, err := koreader.FetchStatsFromEnv()
	if err != nil {
		log.Printf("Warning: Failed to fetch book statistics: %v", err)
		readingStats = &koreader.ReadingStats{
			CurrentBooks:  []koreader.BookStats{},
			FinishedBooks: []koreader.BookStats{},
		}
	} else {
		log.Printf("✓ Fetched statistics for %d books", readingStats.TotalBooks)
	}

	data := struct {
		ReadingStats *koreader.ReadingStats
	}{
		ReadingStats: readingStats,
	}
	templates.ExecuteTemplate(w, "books.html", data)
}
