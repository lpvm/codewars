package main

// ascii 65 -> A  90 -> Z
// ascii 0 -> 48  9 -> 57

// PlayPass is a function
func PlayPass(s string, n int) string {
	runes := []rune(s)
	transf123 := []rune{}
	intl := 0
	minCLetter := 65
	maxCLetter := 90
	minLLetter := 97
	maxLLetter := 122
	minNr := 48
	maxNr := 57
	for _, l := range runes {
		intl = int(l)
		switch {
		// 1. shift each letter by number with circular shift if needed
		case intl >= minCLetter && intl <= maxCLetter:
			{
				n := intl + n
				if n > maxCLetter {
					n = n - maxCLetter + minCLetter - 1
				}
				transf123 = append(transf123, rune(n))
			}
		// 2. replace digit with complement
		case intl >= minNr && intl <= maxNr:
			transf123 = append(transf123, rune(9-(intl-48)+48))
		// 3. keep non alphabetic and non digit chars
		default:
			transf123 = append(transf123, rune(intl))
		}
	}
	transf4 := []rune{}
	// 4. downcase and upcase
	for i, l := range transf123 {
		intl = int(l)
		switch {
		// is an upcase letter
		case i%2 == 1 && intl >= minCLetter && intl <= maxCLetter:
			transf4 = append(transf4, rune(intl-65+97))
		case (i%2 == 0) && intl >= minLLetter && intl <= maxLLetter:
			transf4 = append(transf4, rune(intl-97+65))
		default:
			transf4 = append(transf4, l)
		}
	}
	// 5. Reverse
	transf5 := []rune{}
	for i := len(transf4) - 1; i >= 0; i-- {
		transf5 = append(transf5, transf4[i])
	}

	return string(transf5)
}

func main() {
	//r := PlayPass("BORN IN 2015!", 1)
	// 	orig := []rune{56, 77, 33, 72, 73, 32, 32, 50, 70, 50, 83, 78, 52, 78, 68, 32, 70, 33, 52, 89, 86, 69, 90, 52, 74, 32, 53, 68, 87, 33, 72, 50, 36, 72, 80, 32, 69, 71, 71, 52, 67, 33, 51, 50, 36}
	origs := "8M!HI  2F2SN4ND F!4YVEZ4J 5DW!H2$HP EGG4C!32$"
	origsresp := PlayPass(origs, 5)
}

// orig:  [56 77 33 72 73 32 32 50 70 50 83 78 52 78 68 32 70 33 52 89 86 69 90 52 74 32 53 68 87 33 72 50 36 72 80 32 69 71 71 52 67 33 51 50 36]
// t123:  [49 87 33 82 83 32 32 55 80 55 -61 88 53 88 78 32 80 33 53 -55 -58 79 -54 53 84 32 52 78 -57 33 82 55 36 82 90 32 79 81 81 53 77 33 54 55 36]
// trf4:  [49 119 33 114 83 32 32 55 80 55 -61 120 53 120 78 32 80 33 53 -55 -58 111 -54 53 84 32 52 110 -57 33 82 55 36 114 90 32 79 113 81 53 77 33 54 55 36]

// choose a text in capital letters including or not digits and non alphabetic characters,
//
// 1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
// 2. replace each digit by its complement to 9,
// 3. keep such as non alphabetic and non digit characters,
// 4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
// 5. reverse the whole result.
// #Example:
//
// your text: "BORN IN 2015!", shift 1
// 1 + 2 + 3 -> "CPSO JO 7984!"
// 4 "CpSo jO 7984!"
// 5 "!4897 Oj oSpC"
