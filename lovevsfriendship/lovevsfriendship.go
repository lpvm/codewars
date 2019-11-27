package main

import (
	"fmt"
)

// WordsToMarks is a function
func WordsToMarks(s string) int {
	// a at ascii decimal 97
	sum := 0
	r := []rune(s)
	for _, l := range r {
		sum += int(l) - 96
	}
	return sum
}

func main() {
	w := "love"
	s := WordsToMarks(w)
	fmt.Println(s)
}

// If　a = 1, b = 2, c = 3 ... z = 26
// l + o + v + e = 54
// f + r + i + e + n + d + s + h + i + p = 108
