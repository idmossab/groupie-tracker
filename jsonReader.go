package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func fetchArtist(url string, isSingle bool) interface{} {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if isSingle {
		var artist Artist
		err = json.NewDecoder(resp.Body).Decode(&artist)
		if err != nil {
			log.Print(err)
		}
		return artist
	} else {
		var artists []Artist
		err = json.NewDecoder(resp.Body).Decode(&artists)
		if err != nil {
			log.Print(err)
		}
		return artists
	}
}

func fetchData[T any](url string) T {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result T
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
	}
	return result
}


func fetchArtistWithLocation(id string) Artist {
	artist := fetchData[Artist](ArtistsURL + "/" + id)
	locations := fetchData[Location](LocationsURL+"/"+id) 
	artist.Locations = locations 
	return artist
}