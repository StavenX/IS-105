package main

import (
	"net/http"

	"./1_datanorge"
	"./2_valglokaler"
	"./3_hotell"
	"./4_miljostasjoner"
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