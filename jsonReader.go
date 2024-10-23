package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchArtist(url string) Artist {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artist Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		log.Fatal(err)
	}
	return artist
}
