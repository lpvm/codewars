package main

import "fmt"

//SumCubes function
func SumCubes(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

func main() {
	fmt.Println(SumCubes(1))
	fmt.Println(SumCubes(3))
	fmt.Println(SumCubes(5))
}
