package main

import "net/http"

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorHandler(w, http.StatusMethodNotAllowed)
	}
	if r.URL.Path == "/" {
		artists := fetchData[[]Artist](ArtistsURL)
		renderTemplate(w, "index.html", &artists)
	}else{
		errorHandler(w,http.StatusNotFound)
	}
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	artist := fetchCompleteArtistData(id)
	renderTemplate(w, "detail.html", &artist)
}
