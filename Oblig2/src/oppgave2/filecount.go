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
	//Count the lines
	args := os.Args
	if len(args) == 0 {
		fmt.Println("Missing argument (File)")
		fmt.Println("Usage: go run fileinfo.go myfile.txt")
		return
	}

	filename := args[1]

	file, _ := os.Open(filename)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println(lineCount)

	//Converts file to string
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	str := string(f)

	//Adressing the individual runes in the string
	splitString := []string(strings.Split(str, ""))


	for i := 0; i < len(splitString); i++ {
		MapHandler(splitString[i])
	}
	fmt.Println(len(RuneMap))
	fmt.Println(RuneMap)
}
var array []string

func MapHandler(char string) {
	foundIt := false
	/*for i := 0; i < len(RuneMap); i++ {
		if char == string(RuneMap[rune(i)]) {
			foundIt = true
			fmt.Println("yay")
		}
	}*/
	for i := 0; i < len (array); i++ {
		if char == array[i] {
			foundIt = true
		}
	}
	char2 := []rune(char)
	char3 := int(char2[0])
	if foundIt == false {
		RuneMap[rune(char3)] = 1
		array = append(array, char)
	} else {
		RuneMap[rune(char3)]++
	}
}