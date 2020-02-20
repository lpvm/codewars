package main

import (
	"fmt"
	"strings"
)

// Write a function named first_non_repeating_letter that takes a string input, and returns the first character that is not repeated anywhere in the string.

// For example, if given the input 'stress', the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

// As an added challenge, upper- and lowercase letters are considered the same character, but the function should return the correct case for the initial letter. For example, the input 'sTreSS' should return 'T'.

// If a string contains all repeating characters, it should return an empty string ("") or None -- see sample tests.

// FirstNonRepeatingLetter returns the first character not repeated
func FirstNonRepeatingLetter(s string) string {
	run := []rune(s)
	ls := strings.ToLower(s)

	for _, v := range run {
		if strings.Count(ls, strings.ToLower(string(v))) == 1 {
			return string(v)
		}
	}
	return ""
}

func main() {
	nrchar := FirstNonRepeatingLetter("stress")
	fmt.Println("stress: ", nrchar)
	nrchar = FirstNonRepeatingLetter("sTreSS")
	fmt.Println("sTreSS: ", nrchar)
}
