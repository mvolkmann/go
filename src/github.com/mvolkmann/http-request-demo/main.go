// This uses an iTunes REST service.
// See https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/#lookup.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Album describes a single album.
// json.Unmarshall requires all the fields to be exported (uppercase).
// It matches JSON property names to these case-insensitive.
type Album struct {
	ArtistID         int
	CollectionID     int
	ArtistName       string
	CollectionName   string
	ReleaseDate      string
	PrimaryGenreName string
}

// Albums describes a collection of albums.
// json.Unmarshall requires all the fields to be exported (uppercase).
// It matches JSON property names to these case-insensitive.
type Albums struct {
	ResultCount int
	Results     []Album
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//TODO: Get artist name from command-line argument.
	artist := "Kat Edmonson"
	urlPrefix := "https://itunes.apple.com/search?term="
	getURL := urlPrefix + strings.Replace(artist, " ", "+", -1) + "&entity=album"
	//getURL := url.PathEscape(urlPrefix + artist + "&entity=album")
	//fmt.Println("getURL =", getURL)

	resp, err := http.Get(getURL)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	//fmt.Println("body =", string(body))

	// Using json.Unmarshall is preferred over json.NewDecoder
	// for JSON in HTTP response bodies.  JSON properties that
	// don't match a field in the struct being populated are ignored.
	var albums Albums
	err = json.Unmarshal(body, &albums)
	check(err)
	//fmt.Printf("albums = %+v\n", albums)

	//TODO: Sort albums by release date.
	fmt.Println("Albums by " + artist)
	for _, album := range albums.Results {
		//t, err := time.Parse(time.RFC3339, album.ReleaseDate)
		t, err := time.Parse(time.RFC3339, album.ReleaseDate)
		check(err)
		//fmt.Printf("%s %s\n", album.CollectionName, t.Format("m/d/yyyy"))
		//TODO: Only output name and year?
		fmt.Printf("%s released on %d/%d/%d\n", album.CollectionName, t.Month(), t.Day(), t.Year())
	}
}
