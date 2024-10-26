package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	temp         = template.Must(template.ParseGlob("./templates/*.html"))
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

func main() {

	// Fetch artist data

	http.HandleFunc("/", getHandler)
	http.HandleFunc("/detail", getDetail)
	http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	artists := fetchArtist(ArtistsURL)
	renderTemplate(w,"index.html",artists)
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("id"))
	id:=r.URL.Query().Get("id")
	artists := fetchArtist(ArtistsURL+"/"+id)
	renderTemplate(w,"detail.html",artists)
}
func renderTemplate(w http.ResponseWriter,page string, data interface{}) {
	err := temp.ExecuteTemplate(w, page, data)

	if err != nil {
		http.Error(w, "ERROR SERVER", http.StatusInternalServerError)
	}
}
