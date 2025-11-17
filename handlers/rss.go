package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

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

func RSS(w http.ResponseWriter, r *http.Request) {
	posts, err := loadBlogPostsForRSS()
	if err != nil {
		log.Printf("Error loading blog posts for RSS: %v", err)
		http.Error(w, "Error generating RSS feed", http.StatusInternalServerError)
		return
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].RawDate.After(posts[j].RawDate)
	})

	items := make([]Item, 0, len(posts))
	for _, post := range posts {
		items = append(items, Item{
			Title:       post.Title,
			Link:        "https://aacevski.com/blog/" + post.Slug,
			Description: post.Excerpt,
			PubDate:     post.RawDate.Format(time.RFC1123Z),
			GUID:        "https://aacevski.com/blog/" + post.Slug,
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
			Link:        "https://aacevski.com",
			Description: "breaking code, building tools. software engineer writing about go, typescript, and making things that work.",
			Language:    "en-us",
			AtomLink: AtomLink{
				Href: "https://aacevski.com/rss.xml",
				Rel:  "self",
				Type: "application/rss+xml",
			},
			LastBuildDate: lastBuildDate,
			Items:         items,
		},
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	output, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		log.Printf("Error marshaling RSS feed: %v", err)
		http.Error(w, "Error generating RSS feed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(xml.Header))
	w.Write(output)
}

func loadBlogPostsForRSS() ([]BlogPost, error) {
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
		post, err := parseBlogPostMetadata(filepath.Join(blogDir, file.Name()), slug)
		if err != nil {
			log.Printf("Error parsing blog post %s: %v", file.Name(), err)
			continue
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func parseBlogPostMetadata(filePath, slug string) (BlogPost, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return BlogPost{}, err
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return BlogPost{}, nil
	}

	frontmatter := parts[1]
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

	return post, nil
}
