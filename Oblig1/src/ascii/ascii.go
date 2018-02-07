package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i:=128; i < 256; i++ {
		IterateOverASCIIStringLiteral(i)
	}
	
	text := [16]int{ 34, 32, 8364, 32, 247, 32, 190, 32, 100, 111, 108, 108, 97, 114, 32, 34 }
	ExtendedASCIIText(text)
}

func IterateOverASCIIStringLiteral(i int) {
	toHex := fmt.Sprintf("%X", i)
	toString := string(i)
	toInt64 := int64(i)
	toBinary := strconv.FormatInt(toInt64, 2)
	fmt.Printf("%s %s %s\n", toHex, toString, toBinary)
}

// Kode for Oppgave 2B
func ExtendedASCIIText(text [16]int) {
	var everyLetter string
	for i := 0; i < len(text); i++ {
		everyLetter += string(text[i])
	}

	fmt.Printf("%s", everyLetter)
}
