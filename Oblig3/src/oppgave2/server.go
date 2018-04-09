package main

import (
	"encoding/json"
	"os"
	"log"
	//"io"
	//"net/http"
)

func main() {

	decoder := json.NewDecoder(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := decoder.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}