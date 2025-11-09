package koreader

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"aacevski.com/pkg/utils"
	"github.com/studio-b12/gowebdav"
	_ "modernc.org/sqlite"
)

type BookStats struct {
	Title            string
	Author           string
	Progress         float64
	Pages            int
	CurrentPage      int
	Status           string
	LastRead         time.Time
	TotalReadingTime int
}

type ReadingStats struct {
	CurrentBooks     []BookStats
	FinishedBooks    []BookStats
	TotalBooks       int
	TotalPagesRead   int
	TotalReadingTime int
	BooksThisYear    int
}

type Config struct {
	WebDAVURL string
	Email     string
	Password  string
	DBPath    string
}

func FetchStatsFromEnv() (*ReadingStats, error) {
	config := Config{
		WebDAVURL: utils.GetEnvOrDefault("KOOFR_WEBDAV_URL", "https://app.koofr.net/dav/Koofr"),
		Email:     os.Getenv("KOOFR_EMAIL"),
		Password:  os.Getenv("KOOFR_PASSWORD"),
		DBPath:    utils.GetEnvOrDefault("KOREADER_DB_PATH", "/KOReader/statistics.sqlite3"),
	}

	if config.Email == "" || config.Password == "" {
		return &ReadingStats{CurrentBooks: []BookStats{}, FinishedBooks: []BookStats{}}, nil
	}

	return FetchStats(config)
}

func FetchStats(config Config) (*ReadingStats, error) {
	client := gowebdav.NewClient(config.WebDAVURL, config.Email, config.Password)

	tempDir := os.TempDir()
	dbPath := filepath.Join(tempDir, "koreader.db")

	data, err := client.Read(config.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read database from Koofr: %w", err)
	}

	if err := os.WriteFile(dbPath, data, 0644); err != nil {
		return nil, fmt.Errorf("failed to write temporary database file: %w", err)
	}
	defer os.Remove(dbPath)

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	return parseDatabase(db)
}

func parseDatabase(db *sql.DB) (*ReadingStats, error) {
	bookDir := utils.GetEnvOrDefault("KOREADER_BOOK_DIR", "/books")

	query := `
		SELECT COALESCE(title, 'Unknown Title'), COALESCE(authors, 'Unknown Author'), 
		       COALESCE(pages, 0), COALESCE(percent_finished, 0), 
		       COALESCE(total_time_in_sec, 0), COALESCE(last_open, 0), 
		       COALESCE(directory, '')
		FROM book
		WHERE id NOT IN (1, 2) AND (directory LIKE ? OR directory LIKE ?)
		ORDER BY last_open DESC
	`

	rows, err := db.Query(query, bookDir+"/%", "%"+bookDir+"/%")
	if err != nil {
		return parseAlternativeSchema(db)
	}
	defer rows.Close()

	return processBookRows(rows, db, bookDir, true)
}

func parseAlternativeSchema(db *sql.DB) (*ReadingStats, error) {
	columns := discoverBookTableColumns(db)
	if len(columns) == 0 || !hasColumn(columns, "title") {
		return &ReadingStats{CurrentBooks: []BookStats{}, FinishedBooks: []BookStats{}}, nil
	}

	query, args := buildBookQuery(columns)
	rows, err := db.Query(query, args...)
	if err != nil {
		return &ReadingStats{CurrentBooks: []BookStats{}, FinishedBooks: []BookStats{}}, nil
	}
	defer rows.Close()

	bookDir := utils.GetEnvOrDefault("KOREADER_BOOK_DIR", "/books")
	return processBookRows(rows, db, bookDir, hasColumn(columns, "directory"))
}

func processBookRows(rows *sql.Rows, db *sql.DB, bookDir string, hasDirectory bool) (*ReadingStats, error) {
	stats := &ReadingStats{
		CurrentBooks:  []BookStats{},
		FinishedBooks: []BookStats{},
	}

	currentYear := time.Now().Year()

	for rows.Next() {
		var title, authors, directory string
		var pages int
		var percentFinished float64
		var totalTimeInSec int
		var lastOpen int64

		if hasDirectory {
			if err := rows.Scan(&title, &authors, &pages, &percentFinished, &totalTimeInSec, &lastOpen, &directory); err != nil {
				continue
			}
			if directory != "" && !strings.Contains(directory, bookDir) {
				continue
			}
		} else {
			if err := rows.Scan(&title, &authors, &pages, &lastOpen); err != nil {
				continue
			}
		}

		if lastOpen == 0 {
			continue
		}

		lastRead := time.Unix(lastOpen, 0)
		currentPage, totalTime := fetchBookProgress(db, title)

		status, progress := calculateProgress(pages, currentPage, lastRead, totalTime)

		book := BookStats{
			Title:            title,
			Author:           authors,
			Progress:         progress,
			Pages:            pages,
			CurrentPage:      currentPage,
			Status:           status,
			LastRead:         lastRead,
			TotalReadingTime: totalTime / 60,
		}

		stats.TotalBooks++
		stats.TotalPagesRead += currentPage
		stats.TotalReadingTime += totalTime / 60

		switch status {
		case "finished":
			stats.FinishedBooks = append(stats.FinishedBooks, book)
			if lastRead.Year() == currentYear {
				stats.BooksThisYear++
			}
		case "reading":
			stats.CurrentBooks = append(stats.CurrentBooks, book)
		}
	}

	sortBooks(stats)
	return stats, nil
}

func fetchBookProgress(db *sql.DB, title string) (int, int) {
	var bookID int
	if err := db.QueryRow(`SELECT id FROM book WHERE title = ? LIMIT 1`, title).Scan(&bookID); err != nil {
		return 0, 0
	}

	var currentPage int
	pageQuery := `SELECT page FROM page_stat_data WHERE id_book = ? ORDER BY start_time DESC LIMIT 1`
	if err := db.QueryRow(pageQuery, bookID).Scan(&currentPage); err != nil {
		currentPage = 0
	}

	var totalTime int
	timeQuery := `SELECT COALESCE(SUM(duration), 0) FROM page_stat_data WHERE id_book = ?`
	if err := db.QueryRow(timeQuery, bookID).Scan(&totalTime); err != nil {
		totalTime = 0
	}

	return currentPage, totalTime
}

func calculateProgress(pages, currentPage int, lastRead time.Time, totalTime int) (string, float64) {
	status := "reading"
	progress := 0.0

	if pages > 0 && currentPage > 0 {
		progress = (float64(currentPage) / float64(pages)) * 100
		if currentPage >= pages {
			status = "finished"
			progress = 100
		}
	} else if pages > 0 {
		daysSinceRead := time.Since(lastRead).Hours() / 24
		if daysSinceRead > 60 && totalTime > 0 {
			status = "finished"
			progress = 100
		}
	}

	return status, progress
}

func sortBooks(stats *ReadingStats) {
	sort.Slice(stats.CurrentBooks, func(i, j int) bool {
		return stats.CurrentBooks[i].LastRead.After(stats.CurrentBooks[j].LastRead)
	})
	sort.Slice(stats.FinishedBooks, func(i, j int) bool {
		return stats.FinishedBooks[i].LastRead.After(stats.FinishedBooks[j].LastRead)
	})
}

func discoverBookTableColumns(db *sql.DB) []string {
	rows, err := db.Query("PRAGMA table_info(book)")
	if err != nil {
		return []string{}
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var cid int
		var name, colType string
		var notNull, pk int
		var dfltValue interface{}
		if err := rows.Scan(&cid, &name, &colType, &notNull, &dfltValue, &pk); err != nil {
			continue
		}
		columns = append(columns, name)
	}
	return columns
}

func hasColumn(columns []string, col string) bool {
	for _, c := range columns {
		if c == col {
			return true
		}
	}
	return false
}

func buildBookQuery(columns []string) (string, []interface{}) {
	bookDir := utils.GetEnvOrDefault("KOREADER_BOOK_DIR", "/books")
	hasDir := hasColumn(columns, "directory")

	if hasDir {
		query := `
			SELECT COALESCE(title, 'Unknown Title'), COALESCE(authors, 'Unknown Author'), 
			       COALESCE(pages, 0), 0, 0, COALESCE(last_open, 0), COALESCE(directory, '')
			FROM book
			WHERE id NOT IN (1, 2) AND last_open > 0 AND (directory LIKE ? OR directory LIKE ?)
			ORDER BY last_open DESC LIMIT 50
		`
		return query, []interface{}{bookDir + "/%", "%" + bookDir + "/%"}
	}

	query := `
		SELECT COALESCE(title, 'Unknown Title'), COALESCE(authors, 'Unknown Author'),
		       COALESCE(pages, 0), COALESCE(last_open, 0)
		FROM book
		WHERE id NOT IN (1, 2) AND last_open > 0
		ORDER BY last_open DESC LIMIT 50
	`
	return query, []interface{}{}
}
