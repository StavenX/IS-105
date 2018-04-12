package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"strconv"
	"strconv"
)

type Entries struct {
	Entries []Entry `json:"entries"`
}

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

func main() {
// Open our jsonFile
jsonFile, err := os.Open("valglokaler.json")

// if we os.Open returns an error then handle it
if err != nil {
	fmt.Println(err)
}
fmt.Println("Successfully Opened valglokaler.json")
fmt.Println("https://hotell.difi.no/api/json/kmd/valglokaler/2015/forhand?")

// defer the closing of our jsonFile so that we can parse it later on
defer jsonFile.Close()

// read our opened xmlFile as a byte array.
byteValue, _ := ioutil.ReadAll(jsonFile)


// we initialize our Users array
var entries Entries


// we unmarshal our byteArray which contains our
// jsonFile's content into 'users' which we defined above
json.Unmarshal(byteValue, &entries)


// we iterate through every user within our users array and
// print out the user Type, their name, and their facebook url

for i := 0; i < len(entries.Entries); i++ {
fmt.Println(" ")
fmt.Println("Entry: " + strconv.Itoa(i))
fmt.Println("Kommune: " + entries.Entries[i].Kommune)
fmt.Println("Fylke: " + entries.Entries[i].Fylke)
fmt.Println("Navn: " + entries.Entries[i].Navn)
fmt.Println("Sist endret: " + entries.Entries[i].SistEndret)
fmt.Println("Addresse: " + entries.Entries[i].Addresse)
fmt.Println("Åpningstider: " + entries.Entries[i].Aapningstider)
fmt.Println("Longitude: " + entries.Entries[i].Longitude)
fmt.Println("ID: " + entries.Entries[i].ID)
fmt.Println("Latitude: " + entries.Entries[i].Latitude)
fmt.Println("URL: " + entries.Entries[i].URL)

//fmt.Println(" ")


}}
