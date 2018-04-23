package _hotell

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	getJson()
	http.HandleFunc("/3", PrintToServer)
	http.ListenAndServe(":8080", nil)
}


// ---------------------------------------------------

type Entries struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	By   string `json:"by"`
	Kost   string `json:"kost"`
	Land    string    `json:"land"`
	Makssatser_Natt string `json:"makssatser natt"`
	Verdensdel string `json:"verdensdel"`
	URL string `json:"url"`
}

var entries Entries
var URL = "https://hotell.difi.no/api/json/fad/reise/utland?"

// ---------------------------------------------------

func getJson() {
	url := "https://hotell.difi.no/api/json/fad/reise/utland?"

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
	fmt.Printf("By: %s\n", d.By)
	fmt.Printf("Kost: %s\n", d.Kost)
	fmt.Printf("Land: %s\n", d.Land)
	fmt.Printf("Makssatser_Natt: %s\n", d.Makssatser_Natt)
	fmt.Printf("Verdensdel: %s\n", d.Verdensdel)
	fmt.Printf("URL: %s\n", d.URL)
}

// Prints all the data to server
func PrintToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(entries.Entries); i++ {
		d := entries.Entries[i]
		fmt.Fprintln(writer)
		fmt.Fprintf(writer,"By: %s\n", d.By)
		fmt.Fprintf(writer,"Kost: %s\n", d.Kost)
		fmt.Fprintf(writer,"Land: %s\n", d.Land)
		fmt.Fprintf(writer,"Makssatser_Natt: %s\n", d.Makssatser_Natt)
		fmt.Fprintf(writer,"Verdensdel: %s\n", d.Verdensdel)
		fmt.Fprintf(writer,"URL: %s\n", d.URL)
	}
}