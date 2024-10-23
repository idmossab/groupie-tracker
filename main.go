package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Step 1: Fetch data from the artists API
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Step 2: Read and parse the JSON response for artists
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatal(err)
	}

	// Step 3: For each artist, fetch the locations data and display it
	for _, artist := range artists {
		fmt.Printf("Artist ID: %d\n", artist.ID)
		fmt.Printf("Name: %s\n", artist.Name)
		fmt.Printf("Members: %v\n", artist.Members)
		fmt.Printf("Creation Date: %d\n", artist.CreationDate)
		fmt.Printf("First Album: %s\n", artist.FirstAlbum)
		fmt.Printf("Image URL: %s\n", artist.Image)

		// Fetch the locations data using the LocationsURL
		locationResponse, err := http.Get(artist.LocationsURL)
		if err != nil {
			log.Fatal(err)
		}
		defer locationResponse.Body.Close()

		locationBody, err := ioutil.ReadAll(locationResponse.Body)
		if err != nil {
			log.Fatal(err)
		}

		var locationData Locations
		err = json.Unmarshal(locationBody, &locationData)
		if err != nil {
			log.Fatal(err)
		}

		// Display the locations for the artist
		fmt.Printf("Locations: %v\n", locationData.Locations)
		fmt.Println("----------------------------")
	}
}
