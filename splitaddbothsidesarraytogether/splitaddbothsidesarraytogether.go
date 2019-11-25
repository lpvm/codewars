package main

import (
	"fmt"
)

func sumArr(short, long []int) []int {
	s := []int{}
	longrev := []int{}
	if len(long) > len(short) {
		s = append(s, long[0])
		longrev = long[1:]
	} else {
		longrev = long[:]
	}
	for i := 0; i < len(short); i++ {
		s = append(s, short[i]+longrev[i])
	}
	return s
}

// SplitAndAdd is a function
func SplitAndAdd(numbers []int, n int) []int {
	if len(numbers) <= 1 || n == 0 {
		return numbers
	}

	for {
		n1 := numbers[0 : len(numbers)/2]
		n2 := numbers[len(numbers)/2:]
		numbers = []int{}
		numbers = sumArr(n1, n2)
		n--
		if n == 0 || len(numbers) == 1 {
			break
		}
	}
	return numbers
}

func main() {

	a := SplitAndAdd([]int{1, 2, 3, 4}, 1)
	fmt.Println("end: ", a)
}
