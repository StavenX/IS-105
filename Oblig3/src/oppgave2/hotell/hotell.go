package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type Entries struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	By   string `json:"by"`
	Kost   string `json:"kost"`
	Land    string    `json:"land"`
	Makssatser_Natt string `json:"makssatser natt"`
	Verdensdel string `json:"verdensdel"`
}

func main() {
// Open our jsonFile
jsonFile, err := os.Open("hotell.json")

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
fmt.Println("By: " + entries.Entries[i].By)
fmt.Println("Kost: " + entries.Entries[i].Kost)
fmt.Println("Land: " + entries.Entries[i].Land)
fmt.Println("Makssatser natt: " + entries.Entries[i].Makssatser_Natt)
fmt.Println("Verdensdel: " + entries.Entries[i].Verdensdel)
fmt.Println(" ")
}
}
