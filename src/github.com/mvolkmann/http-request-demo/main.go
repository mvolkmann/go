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
	wrapperType            string
	collectionType         string
	artistId               int
	collectionId           int
	amgArtistId            int
	artistName             string
	collectionName         string
	collectionCensoredName string
	artistViewUrl          string
	collectionViewUrl      string
	artworkUrl60           string
	artworkUrl100          string
	collectionPrice        float32
	collectionExplicitness string
	trackCount             int
	copyright              string
	country                string
	currency               string
	releaseDate            string
	primaryGenreName       string
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
	fmt.Println("body =", string(body))

	//err = json.Unmarshal(body, &albums)
	var albums Albums
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&albums)
	check(err)
	fmt.Printf("albums = %+v\n", albums)
}
