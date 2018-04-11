package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Entries struct {
	Entries []Miljostasjon `json:"entries"`
}
// {"latitude":"58.96403044","navn":"Helgø Mat Stokka","plast":"J","glass_metall":"J","tekstil_sko":"J","longitude":"5.701959134"}
type Miljostasjon struct {
	Latitude   string `json:"latitude"`
	Navn   string `json:"navn"`
	Plast    string    `json:"plast"`
	Glass_Metall string `json:"glass_metall"`
	Tekstil_sko string `json:"tekstil_sko"`
	Longitude string `jsoh:"longitude"`
}

func main() {
// Open our jsonFile
jsonFile, err := os.Open("miljostasjoner.json")

// if we os.Open returns an error then handle it
if err != nil {
	fmt.Println(err)
}
fmt.Println("Successfully Opened hotell.json")
fmt.Println("https://hotell.difi.no/api/json/fad/reise/utland?")

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
// as just an example
for i := 0; i < len(entries.Entries); i++ {
fmt.Println("Entry: " + strconv.Itoa(i))
fmt.Println("Latitude: " + entries.Entries[i].Latitude)
fmt.Println("Navn: " + entries.Entries[i].Navn)
fmt.Println("Plast: " + entries.Entries[i].Plast)
fmt.Println("Glass_Metall: " + entries.Entries[i].Glass_Metall)
fmt.Println("Tekstil_sko: " + entries.Entries[i].Tekstil_sko)
fmt.Println("Longitude: " + entries.Entries[i].Longitude)
fmt.Println(" ")
}
}
