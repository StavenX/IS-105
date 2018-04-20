package _valglokaler

import (
	//"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"strconv"
	"net/http"
	"time"
	//"strings"
)

func main() {
	getJson()
	http.HandleFunc("/2", PrintToServer)
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------

// Structure of all county entries
type Entries struct {
	Entries []Entry `json:"entries"`
}

// Structure of indivdual countey (kommune)
type Entry struct {
	Kommune string `json:"kommune"`
	Fylke string `json:"fylke"`
	Navn string `json:"navn"`
	SistEndret string `json:"sistEndret"`
	Addresse string `json:"addresse"`
	Aapningstider string `json:"aapningstider"`
	Longitude string `json:"lon"`
	ID string `json:"id"`
	Latitude string `json:"lat"`
	URL string `json:"url"`
}

var entries Entries

// ---------------------------------------------------

func getJson() {
	url := "https://hotell.difi.no/api/json/kmd/valglokaler/2015/forhand?"

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

	fmt.Printf("Length: %d\n", len(entries.Entries))

	for i := 0; i < len(entries.Entries); i++ {
		PrintAll(i)
	}
}

// Prints all the data to console
func PrintAll(index int) {
	d:= entries.Entries[index]
	fmt.Println()
	fmt.Printf("Kommune: %s\n", d.Kommune)
	fmt.Printf("Fylke: %s\n", d.Fylke)
	fmt.Printf("Navn: %s\n", d.Navn)
	fmt.Printf("SistEndret: %s\n", d.SistEndret)
	fmt.Printf("Addresse: %s\n", d.Addresse)
	fmt.Printf("Aapningstider: %s\n", d.Aapningstider)
	fmt.Printf("Longitude: %s\n", d.Longitude)
	fmt.Printf("ID: %s\n", d.ID)
	fmt.Printf("Latitude: %s\n", d.Latitude)
	fmt.Printf("URL: %s\n", d.URL)
}

// Prints all the data to server
func PrintToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(entries.Entries); i++ {
		d := entries.Entries[i]
		fmt.Fprintln(writer)
		fmt.Fprintf(writer,"Kommune: %s\n", d.Kommune)
		fmt.Fprintf(writer,"Fylke: %s\n", d.Fylke)
		fmt.Fprintf(writer,"Navn: %s\n", d.Navn)
		fmt.Fprintf(writer,"SistEndret: %s\n", d.SistEndret)
		fmt.Fprintf(writer,"Addresse: %s\n", d.Addresse)
		fmt.Fprintf(writer,"Aapningstider: %s\n", d.Aapningstider)
		fmt.Fprintf(writer,"Longitude: %s\n", d.Longitude)
		fmt.Fprintf(writer,"ID: %s\n", d.ID)
		fmt.Fprintf(writer,"Latitude: %s\n", d.Latitude)
		fmt.Fprintf(writer,"URL: %s\n", d.URL)
	}
}