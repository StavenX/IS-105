package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"strconv"
	"strconv"
	"net/http"
)



func main() {
	openJson()
	http.HandleFunc("/1", printToServer())
	http.ListenAndServe(":8080", nil)
}

// ---------------------------------------------------

type Datasets struct {
	Datasets []Dataset `json:"datasets"`
}

type Dataset struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Antall string `json:"antall"`
	Description []Description `json:"description"`
	Distribution []Distribution `json:"distribution"`
}

type Distribution struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

type Description struct {
	Language string `json:"language"`
	Value string `json:"value"`
}

var datasets Datasets 

// -------------------------------------------------

func openJson() {

}

func printToServer() {

}