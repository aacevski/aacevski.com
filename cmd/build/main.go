package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"andrej.sh/pkg/github"
	"andrej.sh/pkg/koreader"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/joho/godotenv"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

type HomeData struct {
	Name             string
	Nickname         string
	Location         string
	Role             string
	Status           string
	Bio              string
	CurrentlyReading string
	AsciiArt         string
	Work             []WorkItem
	Projects         []ProjectItem
	ReadingStats     *koreader.ReadingStats
	Contributions    *github.ContributionData
}

type WorkItem struct {
	Company     string
	Title       string
	Period      string
	Description string
	URL         string
}

type ProjectItem struct {
	Name        string
	Description string
	URL         string
}

type BlogPost struct {
	Title       string
	Date        string
	RawDate     time.Time
	Slug        string
	Excerpt     string
	Content     template.HTML
	ReadingTime string
}

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	AtomNS  string   `xml:"xmlns:atom,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title         string   `xml:"title"`
	Link          string   `xml:"link"`
	Description   string   `xml:"description"`
	Language      string   `xml:"language"`
	AtomLink      AtomLink `xml:"atom:link"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Items         []Item   `xml:"item"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
}

func main() {
	// Load .env file if it exists (ignore error if file doesn't exist)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	outputDir := "dist"
	if err := os.RemoveAll(outputDir); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("static/templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Fetching book statistics from KOReader...")
	readingStats, err := koreader.FetchStatsFromEnv()
	if err != nil {
		log.Printf("Warning: Failed to fetch book statistics: %v", err)
		readingStats = &koreader.ReadingStats{}
	} else {
		log.Printf("✓ Fetched statistics for %d books", readingStats.TotalBooks)
	}

	data := getHomeData(readingStats)

	indexFile, err := os.Create(filepath.Join(outputDir, "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer indexFile.Close()

	if err := tmpl.Execute(indexFile, data); err != nil {
		log.Fatal(err)
	}

	booksTmpl, err := template.New("books.html").Funcs(template.FuncMap{
		"divf": func(a, b int) float64 {
			return float64(a) / float64(b)
		},
	}).ParseFiles("static/templates/books.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Join(outputDir, "books"), 0755); err != nil {
		log.Fatal(err)
	}

	booksFile, err := os.Create(filepath.Join(outputDir, "books", "index.html"))
	if err != nil {
		log.Fatal(err)
	}
	defer booksFile.Close()

	if err := booksTmpl.Execute(booksFile, data); err != nil {
		log.Fatal(err)
	}

	log.Println("Building blog pages...")
	posts, err := loadBlogPosts()
	if err != nil {
		log.Printf("Warning: Failed to load blog posts: %v", err)
	} else {
		sort.Slice(posts, func(i, j int) bool {
			return posts[i].RawDate.After(posts[j].RawDate)
		})

		blogTmpl, err := template.ParseFiles("static/templates/blog.html")
		if err != nil {
			log.Fatal(err)
		}

		if err := os.MkdirAll(filepath.Join(outputDir, "blog"), 0755); err != nil {
			log.Fatal(err)
		}

		blogListFile, err := os.Create(filepath.Join(outputDir, "blog", "index.html"))
		if err != nil {
			log.Fatal(err)
		}
		defer blogListFile.Close()

		blogData := struct {
			Posts []BlogPost
		}{
			Posts: posts,
		}

		if err := blogTmpl.Execute(blogListFile, blogData); err != nil {
			log.Fatal(err)
		}

		blogPostTmpl, err := template.ParseFiles("static/templates/blog-post.html")
		if err != nil {
			log.Fatal(err)
		}

		for _, post := range posts {
			postDir := filepath.Join(outputDir, "blog", post.Slug)
			if err := os.MkdirAll(postDir, 0755); err != nil {
				log.Fatal(err)
			}

			postFile, err := os.Create(filepath.Join(postDir, "index.html"))
			if err != nil {
				log.Fatal(err)
			}
			defer postFile.Close()

			postData := struct {
				Post BlogPost
			}{
				Post: post,
			}

			if err := blogPostTmpl.Execute(postFile, postData); err != nil {
				log.Fatal(err)
			}
		}

		log.Printf("✓ Built %d blog posts", len(posts))

		log.Println("Building RSS feed...")
		if err := buildRSSFeed(posts, outputDir); err != nil {
			log.Printf("Warning: Failed to build RSS feed: %v", err)
		} else {
			log.Println("✓ Built RSS feed")
		}
	}

	if err := copyDir("static/css", filepath.Join(outputDir, "static/css")); err != nil {
		log.Fatal(err)
	}

	if err := copyDir("static/fonts", filepath.Join(outputDir, "static/fonts")); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("static/favicon.svg", filepath.Join(outputDir, "static/favicon.svg")); err != nil {
		log.Fatal(err)
	}
	if err := copyFile("static/og-image.png", filepath.Join(outputDir, "static/og-image.png")); err != nil {
		log.Fatal(err)
	}
	if err := copyFile("static/robots.txt", filepath.Join(outputDir, "robots.txt")); err != nil {
		log.Fatal(err)
	}

	headersContent := `/*
  X-Frame-Options: DENY
  X-Content-Type-Options: nosniff
  X-XSS-Protection: 1; mode=block
  Referrer-Policy: strict-origin-when-cross-origin

/*.css
  Cache-Control: public, max-age=31536000, immutable

/*.js
  Cache-Control: public, max-age=31536000, immutable

/*.svg
  Cache-Control: public, max-age=31536000, immutable

/*.woff2
  Cache-Control: public, max-age=31536000, immutable
  Access-Control-Allow-Origin: *

/robots.txt
  Content-Type: text/plain
  Cache-Control: public, max-age=3600
`
	if err := os.WriteFile(filepath.Join(outputDir, "_headers"), []byte(headersContent), 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("✓ Static site built successfully in ./dist")
	log.Println("  • index.html (home page)")
	log.Println("  • books/index.html (books page)")
	log.Println("  • blog/index.html (blog listing)")
	log.Println("  • rss.xml (RSS feed)")
	log.Println("  • robots.txt")
	log.Println("  • static/favicon.svg")
	log.Println("  • static/og-image.png")
}

func buildRSSFeed(posts []BlogPost, outputDir string) error {
	items := make([]Item, 0, len(posts))
	for _, post := range posts {
		items = append(items, Item{
			Title:       post.Title,
			Link:        "https://andrej.sh/blog/" + post.Slug,
			Description: post.Excerpt,
			PubDate:     post.RawDate.Format(time.RFC1123Z),
			GUID:        "https://andrej.sh/blog/" + post.Slug,
		})
	}

	lastBuildDate := time.Now().Format(time.RFC1123Z)
	if len(posts) > 0 {
		lastBuildDate = posts[0].RawDate.Format(time.RFC1123Z)
	}

	feed := RSSFeed{
		Version: "2.0",
		AtomNS:  "http://www.w3.org/2005/Atom",
		Channel: Channel{
			Title:       "Andrej Acevski",
			Link:        "https://andrej.sh",
			Description: "breaking code, building tools. software engineer writing about go, typescript, and making things that work.",
			Language:    "en-us",
			AtomLink: AtomLink{
				Href: "https://andrej.sh/rss.xml",
				Rel:  "self",
				Type: "application/rss+xml",
			},
			LastBuildDate: lastBuildDate,
			Items:         items,
		},
	}

	output, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		return err
	}

	rssFile, err := os.Create(filepath.Join(outputDir, "rss.xml"))
	if err != nil {
		return err
	}
	defer rssFile.Close()

	rssFile.Write([]byte(xml.Header))
	rssFile.Write(output)

	return nil
}

func loadBlogPosts() ([]BlogPost, error) {
	var posts []BlogPost
	blogDir := "content/blog"

	files, err := os.ReadDir(blogDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		slug := strings.TrimSuffix(file.Name(), ".md")
		post, err := parseBlogPost(filepath.Join(blogDir, file.Name()), slug)
		if err != nil {
			log.Printf("Error parsing blog post %s: %v", file.Name(), err)
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func parseBlogPost(filePath, slug string) (BlogPost, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return BlogPost{}, err
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return BlogPost{}, fmt.Errorf("invalid frontmatter format")
	}

	frontmatter := parts[1]
	markdown := strings.TrimSpace(parts[2])

	post := BlogPost{Slug: slug}
	lines := strings.Split(frontmatter, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"")

		switch key {
		case "title":
			post.Title = value
		case "date":
			post.Date = value
			if parsedDate, err := time.Parse("2006-01-02", value); err == nil {
				post.RawDate = parsedDate
			}
		case "excerpt":
			post.Excerpt = value
		}
	}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
					html.WithClasses(true),
				),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkhtml.WithHardWraps(),
			goldmarkhtml.WithXHTML(),
			goldmarkhtml.WithUnsafe(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return BlogPost{}, err
	}

	post.Content = template.HTML(buf.String())

	wordCount := len(strings.Fields(markdown))
	readingMinutes := max(wordCount/200, 1)
	post.ReadingTime = fmt.Sprintf("%d min read", readingMinutes)

	return post, nil
}

func getHomeData(readingStats *koreader.ReadingStats) HomeData {
	// Fetch GitHub contributions
	var contributions *github.ContributionData
	ghToken := os.Getenv("GITHUB_TOKEN")
	ghUsername := os.Getenv("GITHUB_USERNAME")
	if ghToken != "" && ghUsername != "" {
		if c, err := github.GetContributions(ghUsername, ghToken); err == nil {
			contributions = c
			log.Printf("✓ Fetched GitHub contributions: %d total", c.TotalContributions)
		} else {
			log.Printf("⚠ Failed to fetch GitHub contributions: %v", err)
		}
	} else {
		log.Println("⚠ GitHub credentials not provided, skipping contribution graph")
	}

	return HomeData{
		Name:             "Andrej Acevski",
		Nickname:         "andrej's shell",
		Location:         "skopje, mk",
		Role:             "eng @ codechem",
		Status:           "breaking code",
		Bio:              "software engineer, open source advocate. fcse graduate. building kaneo and tools that make developers' lives easier.",
		CurrentlyReading: "the art of doing science and engineering",
		AsciiArt: `⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣠⣴⣶⣶⣶⣶⣶⣶⣶⣤⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣸⣿⣿⣿⣿⠓⠻⠿⣿⡻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢻⣿⣿⣿⣿⠀⢀⣤⣀⣤⡄⠀⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣸⣿⣿⡏⣿⣶⣿⣖⣿⡉⠀⠀⢸⣿⡇⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢻⣿⣿⣷⣿⡛⠯⣟⡿⠋⠛⠛⢿⣿⠗⠺⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠛⢿⣿⣿⣿⠙⢷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⢶⣦⣅⣀⡀⠀⢸⡟⢀⣒⣿⣿⣿⣿⣿⣿⣿⣿⣿⠕⢻⣸⣿⣿⡯⠀⠀⢻⡄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣆⠈⠉⠛⣛⣿⣿⣿⠿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣟⢳⣼⣿⣿⣿⡃⠀⠈⠀⡹⡆⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢸⣿⣿⡏⢧⠉⣻⠀⠀⢸⠇⠹⢿⣯⣭⣅⣿⣿⣿⣿⣿⣿⣿⣿⣇⡈⢻⣿⣿⣿⣏⠀⠀⠉⢻⢻⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠸⣿⣿⣿⢀⡼⡿⠀⠀⠘⢧⡀⠈⠙⠛⠓⢿⣿⣿⣿⣿⣿⣿⣿⠁⣰⣿⣿⣿⣿⣿⢆⢷⣦⠀⣼⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣿⣿⣿⣄⣰⠁⠀⠀⠀⠀⠙⢦⡀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⡿⡆⣩⣾⡿⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣼⢿⡷⢿⣿⠷⣄⡀⠀⠀⠀⠀⠀⠱⣄⠀⢸⡇⣿⣿⣿⣿⣿⡟⠀⢹⣿⣿⣿⣿⡿⣣⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣼⡇⢸⡇⠘⣿⣷⡈⠦⣄⡀⠀⠀⠀⠀⠈⠀⣸⠀⣿⣿⣿⡿⢹⡇⠀⠀⢹⡿⢿⣿⣷⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⣸⡟⠀⠘⣧⣴⣹⣿⣷⣄⠺⠭⡙⠖⠀⠀⠀⢀⣨⣾⢿⣿⣿⠃⠘⠁⠀⠀⣸⣷⣿⣿⣿⣿⣿⣟⣿⣿⡆⠀⠀⠀⠀⠀⠀
⠀⠀⢰⣿⠁⠀⠀⢠⠇⣇⠻⣿⡿⣦⣤⣤⣤⣤⣶⣿⡟⠋⠁⠀⣿⡟⠀⠀⢀⣠⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⠀
⠀⠀⣼⣿⣇⠠⠀⡀⠀⡇⠀⢻⣿⣿⣼⡀⠈⢹⠋⠀⠀⠀⣠⠀⣿⣡⣤⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⠀
⠀⢸⣿⣿⣿⣧⣰⣧⠡⢿⠀⢸⣿⣿⢻⢷⣰⡏⠀⢀⣠⣤⣷⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⡄⠀⠀⠀
⠀⣼⣿⣿⣿⣿⣿⡟⢦⣤⣿⣿⣯⣿⣾⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣧⠀⠀⠀
⠀⣿⣿⣿⣿⣿⣿⣿⣦⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀
⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡆⠀
⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢁⣿⣿⣿⣿⣿⣿⣿⣭⣽⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⠀
⢰⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇
⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷
⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿
⣿⣿⣿⡻⣿⡿⠋⠉⠟⡿⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇`,
		Work: []WorkItem{
			{
				Company:     "CodeChem",
				Title:       "software engineer",
				Period:      "'20 - present",
				Description: "building software solutions and working on interesting problems. go, typescript, and whatever gets the job done.",
				URL:         "https://codechem.com",
			},
			{
				Company:     "Kaneo",
				Title:       "founder & engineer",
				Period:      "'24 - present",
				Description: "building an open source project management platform focused on simplicity and efficiency. 2.4k+ stars on github. go, typescript, postgres. making pm tools that don't suck.",
				URL:         "https://kaneo.app",
			},
		},
		Projects: []ProjectItem{
			{
				Name:        "kaneo",
				Description: "open source project management platform. 2.4k+ stars. self-host it, customize it, make it yours. built with typescript and a lot of coffee.",
				URL:         "https://github.com/usekaneo/kaneo",
			},
			{
				Name:        "drim",
				Description: "cli tool to easily deploy your kaneo instance. because deployment should be simple. written in go.",
				URL:         "https://github.com/usekaneo/drim",
			},
			{
				Name:        "andrej.sh",
				Description: "this website. minimal portfolio built with go templates and astro. no unnecessary javascript. fast and clean.",
				URL:         "https://github.com/aacevski/andrej.sh",
			},
		},
		ReadingStats:  readingStats,
		Contributions: contributions,
	}
}

func copyDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0755); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}
