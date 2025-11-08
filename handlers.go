package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/gorilla/mux"
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
	Title   string
	Date    string
	RawDate time.Time
	Slug    string
	Excerpt string
	Content template.HTML
}

type BlogListData struct {
	Posts []BlogPost
}

type BlogPostData struct {
	Post BlogPost
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
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
				Period:      "'25 - present",
				Description: "building an open source project management platform focused on simplicity and efficiency. 2.4k+ stars on github. go, typescript, postgres. making pm tools that don't suck.",
				URL:         "https://usekaneo.com",
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
				Name:        "aacevski.com",
				Description: "this website. minimal portfolio built with go templates and astro. no unnecessary javascript. fast and clean.",
				URL:         "https://github.com/aacevski/aacevski.com",
			},
		},
	}

	templates.ExecuteTemplate(w, "index.html", data)
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := loadBlogPosts()
	if err != nil {
		log.Printf("Error loading blog posts: %v", err)
		http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].RawDate.After(posts[j].RawDate)
	})

	data := BlogListData{
		Posts: posts,
	}
	templates.ExecuteTemplate(w, "blog.html", data)
}

func BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, err := loadBlogPost(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := BlogPostData{
		Post: post,
	}
	templates.ExecuteTemplate(w, "blog-post.html", data)
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

func loadBlogPost(slug string) (BlogPost, error) {
	filePath := filepath.Join("content/blog", slug+".md")
	return parseBlogPost(filePath, slug)
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
	return post, nil
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books coming soon..."))
}

func RSSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/rss+xml")
	w.Write([]byte("RSS feed coming soon..."))
}
