package main

import (
	"net/http"

	"html/template"
	"./pages/1_datanorge"
	"./pages/2_valglokaler"
	"./pages/3_hotell"
	"./pages/4_miljostasjoner"
	//"./pages/5_organisasjon"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", helloClient)
	http.HandleFunc("/1", printPage1)
	http.HandleFunc("/2", printPage2)
	http.HandleFunc("/3", printPage3)
	http.HandleFunc("/4", printPage4)
	//http.HandleFunc("/5", Page5)
	http.ListenAndServe(":8080", nil)
}

//for oppgave1
func helloClient(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(writer, "Hello pro client!")
}

// Struct that defines the information that is supposed to be stored
// in the template. This is just for presenting the information.
type Template struct {
	Title string

	Name0 string
	Name1 string
	Name2 string
	Name3 string
	Name4 string
	Name5 string
	Name6 string
	Name7 string

	Value0 string
	Value1 string
	Value2 string
	Value3 string
	Value4 string
	Value5 string
	Value6 string
	Value7 string
}

// Returns the json to the Unmarshal function. Works similarily to the getJson()
// in the respective pages's go file.
func getJson(url string) []byte {
	client := http.Client{
		Timeout: time.Second *2,
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

//adds values to the Template type to make it ready for the html template
func loadTemplate(title string, names, values []string) *Template {
	//adds spacing between name and value for template
	//as well as filling in for missing or empty data
	for i := 0; i < len(names); i++ {
		names[i] += ": "
		if len(values[i]) == 0 {
			values[i] += "-missing-"
		}
	}

	//initialises an empty Template type, needs to happen within function or it will not clear properly
	t := Template{Title:title}
	t.Name0 = names[0]
	t.Value0 = values[0]
	//adds names and values to the Template type based on the names slice length
	for i := 0; i < len(names); i++ {
		switch i {
		case 1:
			t.Name1 = names[i]
			t.Value1 = values[i]
		case 2:
			t.Name2 = names[i]
			t.Value2 = values[i]
		case 3:
			t.Name3 = names[i]
			t.Value3 = values[i]
		case 4:
			t.Name4 = names[i]
			t.Value4 = values[i]
		case 5:
			t.Name5 = names[i]
			t.Value5 = values[i]
		case 6:
			t.Name6 = names[i]
			t.Value6 = values[i]
		case 7:
			t.Name7 = names[i]
			t.Value7 = values[i]
		}
	}
	return &t
}

// Function that writes a header for the HTML page.
func writeTitle (writer http.ResponseWriter, url string) {
	fmt.Fprintln(writer, "<h1 style =\" font-size:3em\"< Datasets from: </h1)")
	fmt.Fprintf(writer, "<a style=\"font-size:2em\" href=\"%s\" target=\"_blank\">%s</a>\n",url, url)
	fmt.Fprint(writer, "<br><br>")
}

// Renders the template as a html page
func useTemplate(w http.ResponseWriter, page *Template) {
	template, _ := template.ParseFiles("page-template.html")
	template.Execute(w, page)
}

// ------------------------------------------------------------------
// Pages code below!												|
// ------------------------------------------------------------------

// Prints the first page of datasets.
func printPage1(writer http.ResponseWriter, r *http.Request) {
	// Initialises the page
	page  := _datanorge.Datasets{}

	// Puts json data into the readable format defined in page
	jsonErr := json.Unmarshal(getJson(_datanorge.URL), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "Datasett over forskjellige norske API'er"

	// Writes header to the html page
	writeTitle(writer, _datanorge.URL)

	//Uses template on each object in the json data
	for i := 0; i < len(page.Datasets); i++ {
		names := []string{"ID", "Title", "Antall", "Description"}
		v := page.Datasets[i]
		values := []string{v.ID, v.Title, v.Antall, v.Description[0].Value}
		useTemplate(writer, loadTemplate(title, names, values))
	}
}

// Prints the second page of valglokaler.
func printPage2(writer http.ResponseWriter, r *http.Request) {
	// Initialises the page
	page  := _valglokaler.Entries{}

	// Puts json data into the readable format defined in page
	jsonErr := json.Unmarshal(getJson(_valglokaler.URL), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "Oversikt over forskjellige valglokaler"

	// Writes header to the html page
	writeTitle(writer, _valglokaler.URL)

	//Uses template on each object in the json data
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Kommune", "Fylke", "Navn", "SistEndret", "Addresse", "Aapningstider",
		"Longitude,", "ID", "Latitude", "URL"}
		v := page.Entries[i]
		values := []string{v.Kommune, v.Fylke, v.Navn, v.SistEndret, v.Addresse,
		v.Aapningstider, v.Longitude, v.ID, v.Latitude, v.URL}

		useTemplate(writer, loadTemplate(title, names, values))
	}
}

// Prints the third page of hotell.
func printPage3(writer http.ResponseWriter, r *http.Request) {
	// Initialises the page
	page  := _hotell.Entries{}

	// Puts json data into the readable format defined in page
	jsonErr := json.Unmarshal(getJson(_hotell.URL), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "Oversikt over hotell i utlandet"

	// Writes header to the html page
	writeTitle(writer, _hotell.URL)

	//Uses template on each object in the json data
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"By", "Kost", "Land", "Makssatser_natt", "Verdensdel"}
		v := page.Entries[i]
		values := []string{v.By, v.Kost, v.Land, v.Makssatser_Natt, v.Verdensdel}

		useTemplate(writer, loadTemplate(title, names, values))
	}
}

// Prints the fouth page of miljøstasjoner.
func printPage4(writer http.ResponseWriter, r *http.Request) {
	// Initialises the page
	page  := _miljostasjoner.Entries{}

	// Puts json data into the readable format defined in page
	jsonErr := json.Unmarshal(getJson(_miljostasjoner.URL), &page)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	title := "Oversikt over miljøstasjoner"

	// Writes header to the html page
	writeTitle(writer, _miljostasjoner.URL)

	//Uses template on each object in the json data
	for i := 0; i < len(page.Entries); i++ {
		names := []string{"Latitude", "Navn", "Plast", "Glass_metall", "Tekstil_Sko",
		"Longitude"}
		v := page.Entries[i]
		values := []string{v.Latitude, v.Navn, v.Plast, v.Glass_Metall, v.Tekstil_sko,
		v.Longitude}

		useTemplate(writer, loadTemplate(title, names, values))
	}
}