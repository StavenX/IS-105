package _miljostasjoner

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	getJson()
	http.HandleFunc("/4", PrintToServer)
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------

type Entries struct {
	Entries []Miljostasjon `json:"entries"`
}

type Miljostasjon struct {
	Latitude   string `json:"latitude"`
	Navn   string `json:"navn"`
	Plast    string    `json:"plast"`
	Glass_Metall string `json:"glass_metall"`
	Tekstil_sko string `json:"tekstil_sko"`
	Longitude string `jsoh:"longitude"`
}

var entries Entries
var URL = "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

// ---------------------------------------------------

func getJson() {
	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

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

func PrintAll(index int) {
	d:= entries.Entries[index]
	fmt.Println()
	fmt.Printf("Latitude: %s\n", d.Latitude)
	fmt.Printf("Navn: %s\n", d.Navn)
	fmt.Printf("Plast: %s\n", d.Plast)
	fmt.Printf("Glass_Metall: %s\n", d.Glass_Metall)
	fmt.Printf("Tekstil_sko: %s\n", d.Tekstil_sko)
	fmt.Printf("Longitude: %s\n", d.Longitude)
}

func PrintToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(entries.Entries); i++ {
		d := entries.Entries[i]
		fmt.Fprintln(writer)
		fmt.Fprintf(writer,"Latitude: %s\n", d.Latitude)
		fmt.Fprintf(writer,"Navn: %s\n", d.Navn)
		fmt.Fprintf(writer,"Plast: %s\n", d.Plast)
		fmt.Fprintf(writer,"Glass_Metall: %s\n", d.Glass_Metall)
		fmt.Fprintf(writer,"Tekstil_sko: %s\n", d.Tekstil_sko)
		fmt.Fprintf(writer,"Longitude: %s\n", d.Longitude)
	}
}