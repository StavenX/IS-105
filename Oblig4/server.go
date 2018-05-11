package main

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
)

// ----------------------------

// Struct for the whole array of Earthquakes
type Earthquakes struct {
	Earthquakes []Earthquake `json:"features"`
}

// Struct for the Earthquake data itself
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

// Struct for the header that exists in each individual Earthquake json
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

var entries Earthquakes
var header Header

// ----------------------------

// Excecuting things happens here
func main() {
 	openServer()
}

// Uses an URL as parameter for the getJson. Works only on URL's from
// https://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
func GetJson(_url string) error {

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

	return jsonerr
}

// Opens a server on the given port. Will greet the user on a default path
func openServer() {

	/* Handles every "/" request. Also works on other "/" that are not handled.*/
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", printHello)
	//http.HandleFunc("/map", test)
	//http.HandleFunc("/smth/", smthHandler)

	// Handler for the individual pages we have selected.
	http.HandleFunc("/1", 	PrintEarthquakesToServerAllHOUR)
	http.HandleFunc("/2", 	PrintEarthquakesToServerAllDAY)
	http.HandleFunc("/3", 	PrintEarthquakesToServerAllWeek)
	http.HandleFunc("/4", 	PrintEarthquakesToServerAllMonth)

	// Opens the server on the given port
	http.ListenAndServe(":8080", nil)
}

func printHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello");
}

func test(writer http.ResponseWriter, request *http.Request) {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./public"))))
}

/*
func smthHandler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/") || (r.URL.Path != "/1") || (r.URL.Path != "/2") || (r.URL.Path != "/3") || (r.URL.Path != "/4") {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Error")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
*/

// Takes Unix time as parameter, and returns time in a readable format
func getUnixAsReadable(_time int64) time.Time  {
	t := time.Unix((
		int64(_time) / 1000), // +
		//int64(_timeZone) * 60,
		0)
	return t
}

// Prints the JSON header to the server
func PrintHeaderToServer(writer http.ResponseWriter, request *http.Request) {
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

	// Print for BBOX (coordination data)
	fmt.Fprintln(writer, "Minimum Longitude: ", header.Bbox[0])
	fmt.Fprintln(writer, "Minimum Latitude: ", header.Bbox[1])
	fmt.Fprintln(writer, "Minimum Depth: ", header.Bbox[2])
	fmt.Fprintln(writer, "Maximum Longitude: ", header.Bbox[3])
	fmt.Fprintln(writer, "Maximum Latitude: ", header.Bbox[4])
	fmt.Fprintln(writer, "Maximum Depth: ", header.Bbox[5])
	fmt.Fprintln(writer)
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
}

// Function to print individual information about an earthquake. Reduces code duplucation, and prevents
// the need for writing many, long functions with the same data
func PrintEarthquakeInformation(writer http.ResponseWriter) {

	fmt.Fprintln(writer)
	fmt.Fprintln(writer, "List of earthquakes: ")

	for i := 0; i < len(entries.Earthquakes); i++ {
		d:= entries.Earthquakes[i]
		fmt.Fprintln(writer)
		fmt.Fprintln(writer,"Place: ", 			d.Properties.Place)
		fmt.Fprintln(writer,"Magnitude: ", 		d.Properties.Magnitude)
		fmt.Fprintln(writer,"Time: ", 			getUnixAsReadable(d.Properties.Time))
		fmt.Fprintln(writer,"Updated: ", 		getUnixAsReadable(d.Properties.Updated))
		fmt.Fprintln(writer,"TimeZone: ", 		d.Properties.TimeZone)
		fmt.Fprintln(writer,"Detail: ", 			d.Properties.Detail)
		fmt.Fprintln(writer,"Felt: ", 			d.Properties.Felt)
		fmt.Fprintln(writer,"CDI: ", 			d.Properties.CDI)
		fmt.Fprintln(writer,"MMI: ", 			d.Properties.MMI)
		fmt.Fprintln(writer,"Alert: ", 			d.Properties.Alert)
		fmt.Fprintln(writer,"Status: ", 			d.Properties.Status)
		fmt.Fprintln(writer,"Tsunami: ", 		d.Properties.Tsunami)
		fmt.Fprintln(writer,"Significance: ",	d.Properties.SIG)
		fmt.Fprintln(writer,"NET: ", 			d.Properties.NET)
		fmt.Fprintln(writer,"Code: ", 			d.Properties.Code)
		fmt.Fprintln(writer,"IDS: ", 			d.Properties.IDS)
		fmt.Fprintln(writer,"Sources: ", 		d.Properties.Sources)
		fmt.Fprintln(writer,"Types: ", 			d.Properties.Types)
		fmt.Fprintln(writer,"NST: ", 			d.Properties.NST)
		fmt.Fprintln(writer,"DMIN: ", 			d.Properties.DMIN)
		fmt.Fprintln(writer,"RMS: ", 			d.Properties.RMS)
		fmt.Fprintln(writer,"GAP: ", 			d.Properties.GAP)
		fmt.Fprintln(writer,"MagType: ", 		d.Properties.MagType)
		fmt.Fprintln(writer,"Type: ", 			d.Properties.Type)
		fmt.Fprintln(writer,"LongLatDepth: ", 	header.Features[i].Geometry.Coordinates)
	}
}

// Prints earthquake data to local console
func PrintEarthquakesToConsole() {

	fmt.Println()
	fmt.Println("List of earthquakes: ")

	for i := 0; i < len(entries.Earthquakes); i++ {
		d:= entries.Earthquakes[i]
		fmt.Println()
		fmt.Println("Magnitude: ", 		d.Properties.Magnitude)
		fmt.Println("Place: ", 			d.Properties.Place)
		fmt.Println("Time: ", 			d.Properties.Time)
		fmt.Println("Updated: ", 		d.Properties.Updated)
		fmt.Println("TimeZone: ", 		d.Properties.TimeZone)
		fmt.Println("Detail: ", 			d.Properties.Detail)
		fmt.Println("Felt: ", 			d.Properties.Felt)
		fmt.Println("CDI: ", 			d.Properties.CDI)
		fmt.Println("MMI: ", 			d.Properties.MMI)
		fmt.Println("Alert: ", 			d.Properties.Alert)
		fmt.Println("Status: ", 			d.Properties.Status)
		fmt.Println("Tsunami: ", 		d.Properties.Tsunami)
		fmt.Println("Significance: ",	d.Properties.SIG)
		fmt.Println("NET: ", 			d.Properties.NET)
		fmt.Println("Code: ", 			d.Properties.Code)
		fmt.Println("IDS: ", 			d.Properties.IDS)
		fmt.Println("Sources: ", 		d.Properties.Sources)
		fmt.Println("Types: ", 			d.Properties.Types)
		fmt.Println("NST: ", 			d.Properties.NST)
		fmt.Println("DMIN: ", 			d.Properties.DMIN)
		fmt.Println("RMS: ", 			d.Properties.RMS)
		fmt.Println("GAP: ", 			d.Properties.GAP)
		fmt.Println("MagType: ", 		d.Properties.MagType)
		fmt.Println("Type: ", 			d.Properties.Type)

	}
}

// Print earthquake data to the server console
func PrintEarthquakesToServerAllHOUR(writer http.ResponseWriter, request *http.Request,) {
	GetJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_hour.geojson")
	PrintHeaderToServer(writer,request)
	PrintEarthquakeInformation(writer)
}

// Print earthquake data to the server console
func PrintEarthquakesToServerAllDAY(writer http.ResponseWriter, request *http.Request,) {
	GetJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_day.geojson")
	PrintHeaderToServer(writer,request)
	PrintEarthquakeInformation(writer)
}

// Print earthquake data to the server console
func PrintEarthquakesToServerAllWeek(writer http.ResponseWriter, request *http.Request,) {
	GetJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_week.geojson")
	PrintHeaderToServer(writer,request)
	PrintEarthquakeInformation(writer)
}

// Print earthquake data to the server console
func PrintEarthquakesToServerAllMonth(writer http.ResponseWriter, request *http.Request,) {
	GetJson("https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/all_month.geojson")
	PrintHeaderToServer(writer,request)
	PrintEarthquakeInformation(writer)
}
