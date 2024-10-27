package main

import "net/http"

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
	}
	if r.URL.Path == "/" {
		artists,_ := fetchData[[]Artist](ArtistsURL)
		renderTemplate(w, "index.html", &artists)
	} else {
		errorHandler(w, http.StatusNotFound)
	}
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
	}
	if r.URL.Path == "/detail" {
		id := r.URL.Query().Get("id")
		if id == "" {
			errorHandler(w, http.StatusBadRequest)
			return
		}
		artist := fetchCompleteArtistData(id)
		if artist.Name == "" {
			errorHandler(w, http.StatusNotFound) // Artist not found or data is empty
			return
		}
		renderTemplate(w, "detail.html", &artist)
	} else {
		errorHandler(w, http.StatusNotFound)
	}
}
