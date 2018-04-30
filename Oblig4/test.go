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

// API key for Google Maps embedding on website
var apiKey = "AIzaSyCZ7HKaCO5AQUQ27f6bEEzCnj6rOAhs_NA"

// ----------------------------
type Earthquakes struct {
	Earthquakes []Earthquake `json:"features"`
}

type Earthquake struct {
	Type string `json:"type"`
	Properties struct {

		Magnitude float32 	`json:"mag"`
		Place string 		`json:"place"`
		Time int64 			`json:"time"`
		Updated int64 		`json:"time"`
		TimeZone int32 		`json:"tz"`
		URL string 			`json:"url"`
		Detail string 		`json:"detail"`
		Felt int32 			`json:"felt"`
		CDI float32 		`json:"cdi"`
		MMI float32 		`json:"mmi"`
		Alert string 		`json:"alert"`
		Status string 		`json:"status"`
		Tsunami int32 		`json:"tsunami"`
		SIG int32 			`json:"sig"`
		NET string 			`json:"net"`
		Code string 		`json:"code"`
		IDS string 			`json:"ids"`
		Sources string 		`json:"sources"`
		Types string 		`json:"types"`
		NST int32 			`json:"nst"`
		DMIN float32 		`json:"dmin"`
		RMS float32 		`json:"rms"`
		GAP float32 		`json:"gap"`
		MagType string 		`json:"magType"`
		Type string 		`json:"type"`
	} `json:"properties"`
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

func PrintHeaderToConsole() {
	fmt.Println("Type: ", 		header.Type)
	
	fmt.Println("Metadata: ")
	fmt.Println("Generated: ", 	header.Metadata.Generated)
	fmt.Println("URL: ", 		header.Metadata.URL)
	fmt.Println("Title: ", 		header.Metadata.Title)
	fmt.Println("Status: ", 		header.Metadata.Status)
	fmt.Println("API: ", 		header.Metadata.API)
	fmt.Println("Count: ", 		header.Metadata.Count)

}

var entries Earthquakes
var header Header

// ----------------------------

func main() {

 	openServer()

	getJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")

}

// Uses an URL as parameter for the getJson. Works only on URL's from
// https://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
func getJson(_url string) {

	url := _url

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

		PrintEarthquakesToConsole()
}

func openServer() {
	/* Handles every "/" request. Also works on other "/"
	// that are not handled.
	http.HandleFunc("/", helloClient)

	// Handler for the individual pages we have selected.
	http.HandleFunc("/1", printPage1)
	http.HandleFunc("/2", printPage2)
	http.HandleFunc("/3", printPage3)
	http.HandleFunc("/4", printPage4)
	http.HandleFunc("/5", printPage5)

	// Opens the server on the given port
	http.ListenAndServe(":8080", nil)
	*/
}

func PrintEarthquakesToConsole() {
	for i := 0; i < len(entries.Earthquakes); i++ {
		d:= entries.Earthquakes[i]
		fmt.Println()
		fmt.Println("Magnitude: ", 	d.Properties.Magnitude)
		fmt.Println("Place: ", 		d.Properties.Place)
		fmt.Println("Time: ", 		d.Properties.Time)
		fmt.Println("Updated: ", 	d.Properties.Updated)
		fmt.Println("TimeZone: ", 	d.Properties.TimeZone)
		fmt.Println("Detail: ", 		d.Properties.Detail)
		fmt.Println("Felt: ", 		d.Properties.Felt)
		fmt.Println("CDI: ", 		d.Properties.CDI)
		fmt.Println("MMI: ", 		d.Properties.MMI)
		fmt.Println("Alert: ", 		d.Properties.Alert)
		fmt.Println("Status: ", 		d.Properties.Status)
		fmt.Println("Tsunami: ", 	d.Properties.Tsunami)
		fmt.Println("Significance: ",d.Properties.SIG)
		fmt.Println("NET: ", 		d.Properties.NET)
		fmt.Println("Code: ", 		d.Properties.Code)
		fmt.Println("IDS: ", 		d.Properties.IDS)
		fmt.Println("Sources: ", 	d.Properties.Sources)
		fmt.Println("Types: ", 		d.Properties.Types)
		fmt.Println("NST: ", 		d.Properties.NST)
		fmt.Println("DMIN: ", 		d.Properties.DMIN)
		fmt.Println("RMS: ", 		d.Properties.RMS)
		fmt.Println("GAP: ", 		d.Properties.GAP)
		fmt.Println("MagType: ", 	d.Properties.MagType)
		fmt.Println("Type: ", 		d.Properties.Type)

	}
}
