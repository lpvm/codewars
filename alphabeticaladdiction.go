package main

import "fmt"

// rune "z" == 122

// AddLetters function
func AddLetters(letters []rune) rune {
	if len(letters) == 0 {
		return 'z'
	}

	nrLetters := 26
	offset := 96
	acum := 0

	for _, v := range letters {
		acum += int(v) - offset
	}

	timesRolled := (acum - 1) / nrLetters
	acum = acum - timesRolled*nrLetters + offset
	runes := []rune(string(acum))
	return runes[0]
}

func main() {
	fmt.Println(string(AddLetters([]rune{})))
	fmt.Println(string(AddLetters([]rune{'a', 'b', 'c'})))
	fmt.Println(string(AddLetters([]rune{'y', 'c', 'b'})))
	fmt.Println(string(AddLetters([]rune{'z'})))
}

// 	AddLetters([]rune{'a', 'b', 'c'}) = 'f'
// 	AddLetters([]rune{'a', 'b'}) = 'c'
// 	AddLetters([]rune{'z'}) = 'z'
// 	AddLetters([]rune{'z', 'a'}) = 'a'
// 	AddLetters([]rune{'y', 'c', 'b'}) = 'd' // notice the letters overflowing
// 	AddLetters([]rune{}) = 'z'
