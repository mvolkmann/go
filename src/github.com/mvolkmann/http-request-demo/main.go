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
type Album struct {
	ArtistID         int
	CollectionID     int
	ArtistName       string
	CollectionName   string
	ReleaseDate      string
	PrimaryGenreName string
}

// Albums describes a collection of albums.
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

	var albums Albums
	err = json.Unmarshal(body, &albums)
	check(err)
	//fmt.Printf("albums = %+v\n", albums)

	fmt.Println("Albums by " + artist)
	for _, album := range albums.Results {
		//t, err := time.Parse(time.RFC3339, album.ReleaseDate)
		t, err := time.Parse(time.RFC3339, album.ReleaseDate)
		check(err)
		//fmt.Printf("%s %s\n", album.CollectionName, t.Format("m/d/yyyy"))
		fmt.Printf("%s released on %d/%d/%d\n", album.CollectionName, t.Month(), t.Day(), t.Year())
	}
}
