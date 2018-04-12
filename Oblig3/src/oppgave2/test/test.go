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
	Entries []Entry `json:"datasets"`
}

type Entry struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	LandingPage string `json:"landingPage"`
	Issued string `json:"issued"`
	Modified string `json:"modified"`
	Language string `json:"language"`
	Publisher string `json:"publisher"`
	Name string `json:"name"`
	MailBox string `json:"mbox"`
	Keyword string `json:"keyword"`
	Distribusjon string `json:"distribusjon"`
	//Title string
}

func main() {
// Open our jsonFile
jsonFile, err := os.Open("test.json")

// if we os.Open returns an error then handle it
if err != nil {
	fmt.Println(err)
}
fmt.Println("Successfully Opened test.json")
fmt.Println("https://data.norge.no/api/dcat/data.json?page=1")

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
fmt.Println("ID: " + entries.Entries[i].ID)
fmt.Println("Title: " + entries.Entries[i].Title)
fmt.Println("Description: " + entries.Entries[i].Description)
fmt.Println("Landing page: " + entries.Entries[i].LandingPage)
fmt.Println("Issued: " + entries.Entries[i].Issued)
fmt.Println("Modified: " + entries.Entries[i].Modified)
fmt.Println("Language: " + entries.Entries[i].Language)
fmt.Println("Publisher: " + entries.Entries[i].Publisher)
fmt.Println("Name: " + entries.Entries[i].Name)
fmt.Println("MailBox: " + entries.Entries[i].MailBox)
fmt.Println("Keyword: " + entries.Entries[i].Keyword)
fmt.Println("Distribusjon: " + entries.Entries[i].Distribusjon)

//fmt.Println(" ")


}}
