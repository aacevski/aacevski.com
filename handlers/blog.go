package handlers

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

func Blog(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	posts, err := loadBlogPosts()
	if err != nil {
		log.Printf("Error loading blog posts: %v", err)
		http.Error(w, "Error loading blog posts", http.StatusInternalServerError)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].RawDate.After(posts[j].RawDate)
	})

	data := struct {
		Posts []BlogPost
	}{
		Posts: posts,
	}
	templates.ExecuteTemplate(w, "blog.html", data)
}

func BlogPostPage(w http.ResponseWriter, r *http.Request, templates *template.Template) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	post, err := loadBlogPost(slug)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Post BlogPost
	}{
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

