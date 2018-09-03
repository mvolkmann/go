package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	resultCount int
	results     []Album
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	artist := "Kat Edmonson"
	urlPrefix := "https://itunes.apple.com/search?term="
	url := urlPrefix + strings.Replace(artist, " ", "+", -1) + "&entity=album"
	fmt.Println("url =", url)
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)
	//fmt.Println("body =", string(body))

	var albums Albums
	err = json.Unmarshal(body, &albums)
	//decoder := json.NewDecoder(resp.Body)
	//err = decoder.Decode(&albums)
	check(err)
	fmt.Printf("albums = %+v\n", albums)
}
