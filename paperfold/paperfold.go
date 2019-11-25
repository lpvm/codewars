package main

import "fmt"

// PaperFold is a func
func PaperFold(ch chan<- int) {
	ch <- 4
}

func pf(s []int, iter int) []int {
	if iter == 0 {
		return s
	}
	term := 1
	rev := []int{}
	for old := len(s) - 1; old >= 0; old-- {
		t := -1
		if s[old] == 0 {
			t = 1
		} else {
			t = 0
		}
		rev = append(rev, t)
	}
	s = append(s, term)
	s = append(s, rev...)
	s = pf(s, iter-1)
	return s
}

func main() {
	//term := 1
	//s := []int{term}
	//r := pf(s, 2)
	// 	fmt.Println("r: ", r)

	c := make(chan int)
	c <- 2
	go PaperFold(c)
	r := <-c
	fmt.Println("r: ", r)
}

// In each iteration of this process, a 1 is placed at the end of the previous iteration's string, then this string is repeated in reverse order, replacing 0 by 1 and vice versa.
// 0 -> 1
// 1 -> 110
// 2 -> 1101100
// 3 -> 110110011100100
// 4 -> 1101100111001001110110001100100
