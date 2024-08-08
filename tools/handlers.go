package tools

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var data PageData
var tmp *template.Template
func init() {
	tmp = template.Must(template.ParseGlob("templates/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	apiURL := "https://groupietrackers.herokuapp.com/api"
	cards, err := FetchArtistData(apiURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error fetching artist data: %v", err)
		return
	}
	data = PageData{Cards: cards}
	if err := tmp.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	tmp.ExecuteTemplate(w, "404.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
	}
	tmp.ExecuteTemplate(w, "aboutus.html", nil)
}

func Bandinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusBadRequest)
	}
	id := strings.TrimPrefix(r.URL.RawQuery, "=id")
	if id == "" {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	uno, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "/400", http.StatusBadRequest)
		return
	}
	tmp.ExecuteTemplate(w, "bandsinfo.html", data.Cards[uno-1])
}
