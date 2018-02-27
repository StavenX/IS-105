package main

import (
	"fmt"
	"os"
	"log"
)

func info() {

	// Assigns path(?) to file and err
	file, err := os.OpenFile("Oblig2/src/files/text.txt", os.O_RDWR,0)

	// If there is an error, log error
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case err != nil:
		// handle the error and return
	case file.Read():
		// it's a directory
	default:
		// it's not a directory
	}

	defer file.Close()
}

func main() {

	fmt.Println("Programmet skal returnere informasjon om en fil")
	info()
}
