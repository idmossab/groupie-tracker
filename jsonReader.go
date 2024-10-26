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

func fetchDate(url string) []Date{
	resp,err :=http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var Dates []Date
	err=json.NewDecoder(resp.Body).Decode(&Dates)
	if err!=nil{
		log.Fatal(err)
	}
	return Dates
}

func fetchRelation(url string) []Relation{
	resp,err :=http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var Relations []Relation
	err=json.NewDecoder(resp.Body).Decode(&Relations)
	if err!=nil{
		log.Fatal(err)
	}
	return Relations
}