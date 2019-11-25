package main

import "fmt"

// PaperFold is a func
func PaperFold(ch chan<- int) {

	v := 1
	i := 1
	for true {
		v = i
		for v%2 == 0 {
			v /= 2
		}
		if v%4 == 1 {
			ch <- 1
		} else {
			ch <- 0
		}
		i++
	}
}

func main() {
	arr := []int{}
	p := make(chan int, 100)
	go PaperFold(p)

	for i := 0; i < 20; i++ {
		arr = append(arr, <-p)
	}
	fmt.Println(arr)
}
