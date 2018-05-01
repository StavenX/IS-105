package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
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
		Updated int64 		`json:"updated"`
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
	Type string `json:"type"`
	Metadata struct {
		Generated int64  `json:"generated"`
		URL       string `json:"url"`
		Title     string `json:"title"`
		Status    int    `json:"status"`
		API       string `json:"api"`
		Count     int    `json:"count"`
	} `json:"metadata"`
	Bbox     []float64 `json:"bbox"`
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
	}

}

// Prints the JSON header to the server (webpage)
func PrintHeaderToServer(writer http.ResponseWriter, request *http.Request) {
	getJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")
	fmt.Fprintln(writer, "----------------------------------------------------------------")
	fmt.Fprintln(writer, "Type: ", header.Type)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "Metadata: ")
	fmt.Fprintln(writer, "Generated: ", getUnixAsReadable(header.Metadata.Generated))
	fmt.Fprintln(writer, "URL: ", header.Metadata.URL)
	fmt.Fprintln(writer, "Title: ", header.Metadata.Title)
	fmt.Fprintln(writer, "Status: ", header.Metadata.Status)
	fmt.Fprintln(writer, "API: ", header.Metadata.API)
	fmt.Fprintln(writer, "Count: ", header.Metadata.Count)
	fmt.Fprintln(writer)

	// Print for BBOX
	fmt.Fprintln(writer, "Minimum Longitude: ", header.Bbox[0])
	fmt.Fprintln(writer, "Minimum Latitude: ", header.Bbox[1])
	fmt.Fprintln(writer, "Minimum Depth: ", header.Bbox[2])
	fmt.Fprintln(writer, "Maximum Longitude: ", header.Bbox[3])
	fmt.Fprintln(writer, "Maximum Latitude: ", header.Bbox[4])
	fmt.Fprintln(writer, "Maximum Depth: ", header.Bbox[5])
	fmt.Fprintln(writer)

	// Print for Geometry
	for i := 0; i < len(header.Features); i++ {
		fmt.Fprintln(writer, "earthquake entry ", i,  header.Features[i])
	}
	fmt.Fprintln(writer, "----------------------------------------------------------------")
}

// Prints the JSON header to the server console. Use for local printing
func PrintHeaderToConsole() {
	fmt.Println("Type: ", 		header.Type)
	fmt.Println("Metadata for session -- ")
	fmt.Println("Generated: ", 	header.Metadata.Generated)
	fmt.Println("URL: ", 		header.Metadata.URL)
	fmt.Println("Title: ", 		header.Metadata.Title)
	fmt.Println("Status: ", 		header.Metadata.Status)
	fmt.Println("API: ", 		header.Metadata.API)
	fmt.Println("Count: ", 		header.Metadata.Count)

	/*bbox test
	fmt.Println()
	fmt.Println("Minimum Longitude: ", header.BBOX.MinLongitude)
	fmt.Println("Minimum Latitude: ", header.BBOX.MinLatitude)
	fmt.Println("Minimum Depth: ", header.BBOX.MinDepth)
	fmt.Println("Maximum Longitude: ", header.BBOX.MaxLongitude)
	fmt.Println("Maximum Latitude: ", header.BBOX.MaxLatitude)
	fmt.Println("Maximum Depth: ", header.BBOX.MaxDepth)
	*/
}

var entries Earthquakes
var header Header

// ----------------------------

func main() {

 	openServer()

	//getJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")
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
		fmt.Println("JSON entries unmarshal error.")
	}

	header_jsonerr := json.Unmarshal(body, &header)
	if header_jsonerr != nil {
		fmt.Println("JSON header unmarshal error.")
	}

	//fmt.Printf("Amount of earthquakes: %d\n", len(entries.Earthquakes))

	// Prints for earthquake information
	PrintHeaderToConsole()
	PrintEarthquakesToConsole()
}

func openServer() {
	// Handles every "/" request. Also works on other "/"
	// that are not handled.
	http.HandleFunc("/", printHello)

	// Handler for the individual pages we have selected.
	http.HandleFunc("/1", PrintEarthquakesToServer1)
	http.HandleFunc("/2", PrintEarthquakesToServer2)
	http.HandleFunc("/header", PrintHeaderToServer)

	// Opens the server on the given port
	http.ListenAndServe(":8080", nil)

}

// Takes Unix time as parameter, and returns time in a readable format
func getUnixAsReadable(_time int64) time.Time  {
	t := time.Unix((
		int64(_time) / 1000), // +
		//int64(_timeZone) * 60,
		0)

	return t
}

// Prints hello to the client
func printHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello client")
}


func PrintEarthquakesToServer1(writer http.ResponseWriter, request *http.Request,) {
	getJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")

	PrintHeaderToServer(writer,request)
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "List of earthquakes: ")

	for i := 0; i < len(entries.Earthquakes); i++ {
		d:= entries.Earthquakes[i]
		fmt.Fprintln(writer)
		fmt.Fprintln(writer,"Place: ", 		d.Properties.Place)
		fmt.Fprintln(writer,"Magnitude: ", 	d.Properties.Magnitude)

		// Converting time in Unix format to readable time
		//t := time.Unix(int64(d.Properties.Time) / 1000, 0)
		//fmt.Fprintln(writer, "Time: ", t)

		fmt.Fprintln(writer, "Time: ", 		getUnixAsReadable(d.Properties.Time))
		fmt.Fprintln(writer,"Updated: ", 	getUnixAsReadable(d.Properties.Updated))
		fmt.Fprintln(writer,"TimeZone: ", 	d.Properties.TimeZone)
		fmt.Fprintln(writer,"Detail: ", 		d.Properties.Detail)
		fmt.Fprintln(writer,"Felt: ", 		d.Properties.Felt)
		fmt.Fprintln(writer,"CDI: ", 		d.Properties.CDI)
		fmt.Fprintln(writer,"MMI: ", 		d.Properties.MMI)
		fmt.Fprintln(writer,"Alert: ", 		d.Properties.Alert)
		fmt.Fprintln(writer,"Status: ", 		d.Properties.Status)
		fmt.Fprintln(writer,"Tsunami: ", 	d.Properties.Tsunami)
		fmt.Fprintln(writer,"Significance: ",d.Properties.SIG)
		fmt.Fprintln(writer,"NET: ", 		d.Properties.NET)
		fmt.Fprintln(writer,"Code: ", 		d.Properties.Code)
		fmt.Fprintln(writer,"IDS: ", 		d.Properties.IDS)
		fmt.Fprintln(writer,"Sources: ", 	d.Properties.Sources)
		fmt.Fprintln(writer,"Types: ", 		d.Properties.Types)
		fmt.Fprintln(writer,"NST: ", 		d.Properties.NST)
		fmt.Fprintln(writer,"DMIN: ", 		d.Properties.DMIN)
		fmt.Fprintln(writer,"RMS: ", 		d.Properties.RMS)
		fmt.Fprintln(writer,"GAP: ", 		d.Properties.GAP)
		fmt.Fprintln(writer,"MagType: ", 	d.Properties.MagType)
		fmt.Fprintln(writer,"Type: ", 		d.Properties.Type)
	}
}
func PrintEarthquakesToServer2(writer http.ResponseWriter, request *http.Request,) {
	getJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")
	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "List of earthquakes: ")

	for i := 0; i < len(entries.Earthquakes); i++ {
		d:= entries.Earthquakes[i]
		fmt.Fprintln(writer)
		fmt.Fprintln(writer,"Magnitude: ", 	d.Properties.Magnitude)
		fmt.Fprintln(writer,"Place: ", 		d.Properties.Place)
		fmt.Fprintln(writer,"Time: ", 		d.Properties.Time)
		fmt.Fprintln(writer,"Updated: ", 	d.Properties.Updated)
		fmt.Fprintln(writer,"TimeZone: ", 	d.Properties.TimeZone)
		fmt.Fprintln(writer,"Detail: ", 		d.Properties.Detail)
		fmt.Fprintln(writer,"Felt: ", 		d.Properties.Felt)
		fmt.Fprintln(writer,"CDI: ", 		d.Properties.CDI)
		fmt.Fprintln(writer,"MMI: ", 		d.Properties.MMI)
		fmt.Fprintln(writer,"Alert: ", 		d.Properties.Alert)
		fmt.Fprintln(writer,"Status: ", 		d.Properties.Status)
		fmt.Fprintln(writer,"Tsunami: ", 	d.Properties.Tsunami)
		fmt.Fprintln(writer,"Significance: ",d.Properties.SIG)
		fmt.Fprintln(writer,"NET: ", 		d.Properties.NET)
		fmt.Fprintln(writer,"Code: ", 		d.Properties.Code)
		fmt.Fprintln(writer,"IDS: ", 		d.Properties.IDS)
		fmt.Fprintln(writer,"Sources: ", 	d.Properties.Sources)
		fmt.Fprintln(writer,"Types: ", 		d.Properties.Types)
		fmt.Fprintln(writer,"NST: ", 		d.Properties.NST)
		fmt.Fprintln(writer,"DMIN: ", 		d.Properties.DMIN)
		fmt.Fprintln(writer,"RMS: ", 		d.Properties.RMS)
		fmt.Fprintln(writer,"GAP: ", 		d.Properties.GAP)
		fmt.Fprintln(writer,"MagType: ", 	d.Properties.MagType)
		fmt.Fprintln(writer,"Type: ", 		d.Properties.Type)
	}
}


func PrintEarthquakesToConsole() {

	fmt.Println()
	fmt.Println("List of earthquakes: ")

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
