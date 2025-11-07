package main

import (
	"net/http"
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
	w.Write([]byte("Blog coming soon..."))
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books coming soon..."))
}

func RSSHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/rss+xml")
	w.Write([]byte("RSS feed coming soon..."))
}
