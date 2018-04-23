package main

import (
	"net/http"

	"html/template"
	"./pages/1_datanorge"
	"./pages/2_valglokaler"
	"./pages/3_hotell"
	"./pages/4_miljostasjoner"
	"./pages/5_organisasjon"
)

func main() {
	//openJson()
	http.HandleFunc("/1", _datanorge.PrintToServer)
	http.HandleFunc("/2", _valglokaler.PrintToServer)
	http.HandleFunc("/3", _hotell.PrintToServer)
	http.HandleFunc("/4", _miljostasjoner.PrintToServer)
	//http.HandleFunc("/5", _package.PrintToServer)

	http.ListenAndServe(":8080", nil)
}