package main

import (
	"os"
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

var RuneMap = make(map[rune]int)

func main() {

	// Declares 'args' as the first argument given on execution.
	// Also checks if no arguments given (hence if the argument has zero characters), and
	// if so asks the user to type another argument (path).
	// + Also gives usage on how to run the file itself.
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Missing argument (File)")
		fmt.Println("Usage: go run fileinfo.go myfile.txt")
		return
	}

	// Sets 'filePath' to the first argument given to terminal.
	filePath := args[1]

	// Scans the file and counts lines
	file, _ := os.Open(filePath)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Information about " + args[1])
	fmt.Printf("Number of lines in file: %d\n", lineCount)
	fmt.Println("Most Common Runes:")

	//Converts file to string
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	str := string(f)

	//Adressing the individual runes in the string
	splitString := []string(strings.Split(str, ""))


	for i := 0; i < len(splitString); i++ {
		MapHandler(splitString[i])
	}

	MapSorter(5)
}
var Foundletters []string

// Finds each character and counts them
func MapHandler(char string) {

	foundIt := false
	// Loop that finds the characters
	for i := 0; i < len (Foundletters); i++ {
		if char == Foundletters[i] {
			foundIt = true
		}
	}
	// Counts the Characters
	char2 := []rune(char)
	char3 := int(char2[0])
	if foundIt == false {
		RuneMap[rune(char3)] = 1
		Foundletters = append(Foundletters, char)
	} else {
		RuneMap[rune(char3)]++
	}
}

func MapSorter (x int) {
	//Loop repeats x times
	for i := 1; i <= x; i++ {
		number := i
		highestCount := 0
		mostUsed := ""

		//Compares every found letter's mentions against the current highestCount
		for i := 0; i < len(Foundletters); i++ {
			runeOfString := []rune(Foundletters[i])
			mentions := RuneMap[runeOfString[0]]

			//Sets new highest count
			if mentions > highestCount {
				highestCount = mentions
				mostUsed = Foundletters[i]
			}
		}
		mostUsedRune := []rune(mostUsed)

		//special condition to better distinguish a space in the stats
		if mostUsed == " " {
			mostUsed = "(space)"
		}

		//prints #i highest number
		fmt.Printf("%d. Rune: '%s' , Times appeared: %d\n", number, mostUsed, highestCount)

		//deletes the current highest number
		delete(RuneMap, mostUsedRune[0])
	}
}