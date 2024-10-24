package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func fetchArtist(url string) []Artist {
	resp, err := http.Get(url)
	var artists []Artist

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &artist)*/
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		log.Fatal(err)
	}
	return artists
}

func fetchLocation(url string) []Location{
	resp,err :=http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var locations []Location
	err=json.NewDecoder(resp.Body).Decode(&locations)
	if err!=nil{
		log.Fatal(err)
	}
	return locations
}
