package algorithms

import "fmt"

func main() {

}

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {
	end := len(list) -1
	for {
		if end == 0 {
		break
		}
	for i := 0; i < len(list) - 1; i++ {
		if list[i] > list[i+1] {
		list[i], list[i+1] = list[i+1], list[i]
		}
	}
	end -= 1
	}
}


// Implementering av Quicksort algoritmen



