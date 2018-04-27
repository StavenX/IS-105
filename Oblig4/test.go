package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	//"strings"
	//"strconv"
)

// ----------------------------
type Earthquakes struct {
	Earthquakes []Earthquake `json:"features"`
}

type Header struct {
	Type     string `json:"type"`
	Metadata struct {
		Generated int64  `json:"generated"`
		URL       string `json:"url"`
		Title     string `json:"title"`
		Status    int    `json:"status"`
		API       string `json:"api"`
		Count     int    `json:"count"`
	} `json:"metadata"`
}

var entries Earthquakes

// ----------------------------

func main() {

	getJson()

}

func getJson() {
	url := "https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson"

	client := http.Client{
		Timeout: time.Second *2,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Request error.")
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Response error.")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Body error.")
	}

	jsonerr := json.Unmarshal(body, &entries)
	if jsonerr != nil {
		fmt.Println("JSON unmarshal error.")
	}

	fmt.Printf("Length: %d\n", len(entries.Earthquakes))

	for i := 0; i < len(entries.Earthquakes); i++ {
		PrintAll(i)
		fmt.Println(i)
	}
}

func PrintAll(index int) {
	d:= entries.Earthquakes[index]
	fmt.Println()
	fmt.Println("Printing METADATA: ")
	fmt.Println("Type: ", d.Type)
	fmt.Println("Generated: ", d.Metadata.Generated)
	fmt.Println("URL: ", d.Metadata.URL)
	fmt.Println("Title: ", d.Metadata.Title)
	fmt.Println("API: ", d.Metadata.API)
	fmt.Println("Count: ", d.Metadata.Count)
	fmt.Println("Status: ", d.Metadata.Status)
}
