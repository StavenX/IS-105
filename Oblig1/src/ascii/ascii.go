package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i:=128; i < 256; i++ {
		IterateOverASCIIStringLiteral(i)
	}
	ExtendedASCIIText()
}

func IterateOverASCIIStringLiteral(i int) {
	h := fmt.Sprintf("%X", i)
	st := string(i)
	n := int64(i)
	bin := strconv.FormatInt(n, 2)
	fmt.Printf("%s %s %s\n", h, st, bin)
}

// Kode for Oppgave 2B
func ExtendedASCIIText() {
	a := string(34)
	b := string(32)
	c := string(8364)
	d := string(247)
	e := string(190)
	f := string(100)
	g := string(111)
	h := string(108)
	i := string(97)
	j := string(114)
	zz := a+b+c+b+d+b+e+b+f+g+h+h+i+j+b+a
	fmt.Printf("%s", zz)
}