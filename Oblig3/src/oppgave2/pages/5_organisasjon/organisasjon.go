package _organisasjon

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	getJson()
	http.HandleFunc("/5", PrintToServer)
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------

type Entries struct {
	Entries	[]Entry	`json:"entries"`
}

type Entry struct {
	Stiftelsesdato		string	`json:"stiftelsesdato"`
	Organisasjonsform	string	`json:"organisasjonsform"`
	Navn				string	`json:"navn"`
	Tlf					string	`json:"tlf"`
	Forretningsadresse	string	`json:"forretningsadr"`
	Poststed			string	`json:"forradrpoststed"`
	Regdato				string	`json:"regdato"`
}

var entries Entries

// ---------------------------------------------------

func getJson() {
	url := "https://hotell.difi.no/api/json/brreg/enhetsregisteret?page=8"

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
	fmt.Printf("Stiftelsesdato: %s\n", d.Stiftelsesdato)
	fmt.Printf("Organisasjonsform: %s\n", d.Organisasjonsform)
	fmt.Printf("Navn: %s\n", d.Navn)
	fmt.Printf("Tlf: %s\n", d.Tlf)
	fmt.Printf("Forretningsadresse: %s\n", d.Forretningsadresse)
	fmt.Printf("Poststed: %s\n", d.Poststed)
	fmt.Printf("Regdato: %s\n", d.Regdato)
}

// Prints all the data to server
func PrintToServer(writer http.ResponseWriter, request *http.Request) {
	for i := 0; i < len(entries.Entries); i++ {
		d := entries.Entries[i]
		fmt.Fprintln(writer)
		fmt.Fprintf(writer,"Stiftelsesdato: %s\n", d.Stiftelsesdato)
		fmt.Fprintf(writer,"Organisasjonsform: %s\n", d.Organisasjonsform)
		fmt.Fprintf(writer,"Navn: %s\n", d.Navn)
		fmt.Fprintf(writer,"Tlf: %s\n", d.Tlf)
		fmt.Fprintf(writer,"Forretningsadresse: %s\n", d.Forretningsadresse)
		fmt.Fprintf(writer,"Poststed: %s\n", d.Poststed)
		fmt.Fprintf(writer,"Regdato: %s\n", d.Regdato)
	}
}