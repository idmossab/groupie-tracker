package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// declaration Variable
var (
	temp = template.Must(template.ParseGlob("./templates/*.html"))
)

const (
	certFile     = "groupie-tracker.pem"
	keyFile      = "groupie-tracker-key.pem"
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/detail", getDetail)

	fmt.Println("Server starting... on https://localhost:8080")
	//http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	http.ListenAndServe(":8080", nil)

}

// Renders the template with the given data
func renderTemplate(w http.ResponseWriter, page string, data any) {
	err := temp.ExecuteTemplate(w, page, data)
	if err != nil {
		errorHandler(w, http.StatusInternalServerError)
	}
}
