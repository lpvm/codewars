package main

import "fmt"

func recmin(i int) int {
	if i == 1 {
		return 0
	}
	return i - 1 + recmin(i-1)
}

func recmax(m, i int) int {
	if i == m {
		return i
	} else if i > m {
		return 0
	}
	return i + recmax(m, i+1)
}

// SuMin is function f
func SuMin(m int) int64 {
	sum := 0
	for i := 1; i <= m; i++ {
		sum += i*m - recmin(i)
	}
	return int64(sum)
}

// SuMax is function g
func SuMax(m int) int64 {
	sum := 0
	for i := 1; i <= m; i++ {
		sum += i*i + recmax(m, i+1)
	}
	return int64(sum)
}

// SumSum is function h
func SumSum(m int) int64 {
	sum := SuMin(m) + SuMax(m)
	return int64(sum)
}

func main() {
	fmt.Println(SuMin(6))
	fmt.Println(SuMax(6))
	fmt.Println(SuMax(45))
	fmt.Println(SumSum(6))
}

// sumin(6) --> 91
// sumin(45) --> 31395
// sumin(999) --> 332833500
// sumin(5000) --> 41679167500
//
// sumax(6) --> 161
// sumax(45) --> 61755
// sumax(999) --> 665167500
// sumax(5000) --> 83345832500
//
// sumsum(6) --> 252
// sumsum(45) --> 93150
// sumsum(999) --> 998001000
// sumsum(5000) --> 125025000000
