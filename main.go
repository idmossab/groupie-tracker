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

	fmt.Println("Server starting... on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


func renderTemplate(w http.ResponseWriter, page string, data any) {
	err := temp.ExecuteTemplate(w, page, data)

	if err != nil {
		errorHandler(w,http.StatusInternalServerError)
	}
}
