package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

// Fetches and decodes JSON data from the URL.
func fetchData[T any](url string) (T,error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	//Solution :the response body is unexpectedly empty or prematurely closed
	if resp.StatusCode != http.StatusOK {
		return *new(T), fmt.Errorf("error: received status code %d", resp.StatusCode)
	}
	var result T
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Print(err)
	}
	return result,nil
}

// Fetches complete artist data including locations, dates, and relations.
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
