package handlers

import (
	"html/template"
	"time"
)

type BlogPost struct {
	Title   string
	Date    string
	RawDate time.Time
	Slug    string
	Excerpt string
	Content template.HTML
}

