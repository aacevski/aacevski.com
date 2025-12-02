package handlers

import (
	"encoding/xml"
	"log"
	"net/http"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc        string  `xml:"loc"`
	LastMod    string  `xml:"lastmod,omitempty"`
	ChangeFreq string  `xml:"changefreq,omitempty"`
	Priority   float64 `xml:"priority,omitempty"`
}

func Sitemap(w http.ResponseWriter, r *http.Request) {
	posts, err := loadBlogPostsForRSS()
	if err != nil {
		log.Printf("Error loading blog posts for sitemap: %v", err)
		http.Error(w, "Error generating sitemap", http.StatusInternalServerError)
		return
	}

	urls := []URL{
		{
			Loc:        "https://andrej.sh/",
			ChangeFreq: "weekly",
			Priority:   1.0,
		},
		{
			Loc:        "https://andrej.sh/blog",
			ChangeFreq: "weekly",
			Priority:   0.9,
		},
		{
			Loc:        "https://andrej.sh/books",
			ChangeFreq: "monthly",
			Priority:   0.8,
		},
	}

	// Add blog posts
	for _, post := range posts {
		urls = append(urls, URL{
			Loc:        "https://andrej.sh/blog/" + post.Slug,
			LastMod:    post.RawDate.Format("2006-01-02"),
			ChangeFreq: "monthly",
			Priority:   0.7,
		})
	}

	sitemap := URLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		log.Printf("Error marshaling sitemap: %v", err)
		http.Error(w, "Error generating sitemap", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(xml.Header))
	w.Write(output)
}

