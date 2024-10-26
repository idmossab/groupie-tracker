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
	artists := fetchData[[]Artist](ArtistsURL)
	renderTemplate(w, "index.html", &artists)
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	id:=r.URL.Query().Get("id")
	fmt.Println(r.URL.Query().Get("id"))
	artist := fetchData[Artist](ArtistsURL + "/"+id)
	renderTemplate(w, "detail.html", &artist)
}
func renderTemplate(w http.ResponseWriter, page string, data any) {
	err := temp.ExecuteTemplate(w, page, data)

	if err != nil {
		http.Error(w, "ERROR SERVER", http.StatusInternalServerError)
	}
}
