package main

import (
	"fmt"
)

// BaumSweet is a function
func BaumSweet(ch chan<- int) {
	d := 0
	arr := []int{}
	for {
		switch {
		case d == 0:
			{
				arr = append(arr, 1)
				ch <- 1
			}
		case d%4 == 0:
			{
				r := arr[d/4]
				arr = append(arr, r)
				ch <- r
			}
		case d%2 == 1:
			{
				r := arr[(d-1)/2]
				arr = append(arr, r)
				ch <- r
			}
		default:
			{
				arr = append(arr, 0)
				ch <- 0
			}
		}
		d++
	}
}

func main() {
	arr := make([]int, 20)
	p := make(chan int, 100)
	go BaumSweet(p)

	for i := 0; i < 20; i++ {
		arr[i] = <-p
	}
	//     []int{1,1,0,1,1,0,0,1,0,1,0,0,1,0,0,1,1,0,0,1}

	fmt.Println("arr", arr)
}

// The Baum–Sweet sequence is an infinite automatic sequence of 0s and 1s defined by the rule:
// bn = 1 if the binary representation of n contains no block of consecutive 0s of odd length;
// bn = 0 otherwise;
// for n ≥ 0.
//
// b_n = b_((n-1)/2) if n is odd;
// b_n = b_(n/4) if n is divisible by 4
// b_n = 0 otherwise.
