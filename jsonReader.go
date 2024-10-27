package main

import (
	"encoding/json"
	"log"
	"net/http"
)

/*func fetchArtist(url string, isSingle bool) interface{} {
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
}*/

func fetchData[T any](url string) (T,error) {
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
	return result,nil
}

func fetchCompleteArtistData(id string) (Artist,bool) {
	artist,err := fetchData[Artist](ArtistsURL + "/" + id)
	// Check for error or if artist is empty
	if err != nil || artist.Name == "" { 
		return Artist{},true
	}
	locations,_ := fetchData[Location](LocationsURL + "/" + id)
	concertDates,_ := fetchData[Date](DatesURL + "/" + id)
	relation,_ := fetchData[Relation](RelationURL + "/" + id)
	artist.Locations = locations
	artist.ConcertDates = concertDates
	artist.Relations = relation
	return artist,false
}
