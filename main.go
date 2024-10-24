package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	temp := template.Must(template.ParseFiles("index.html"))
	// Fetch artist data
	artistURL := "https://groupietrackers.herokuapp.com/api/artists"
	locationURL:="https://groupietrackers.herokuapp.com/api/locations"
	artists := fetchArtist(artistURL)
	locations:=fetchLocation(locationURL)

	for i := range artists {
		for _, loc := range locations {
			if artists[i].ID == loc.ID {
				artists[i].Locations =loc
				break
			}
		}
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err:=temp.Execute(w, artists)
		if err!=nil{
			log.Fatal(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
