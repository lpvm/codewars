// Take 2 strings s1 and s2 including only letters from ato z. Return a new
// sorted string, the longest possible, containing distinct letters,
// each taken only once - coming from s1 or s2

package main

import (
	"fmt"
	"sort"
	"strings"
)

func uniques(s string) []string {
	m := make(map[string]int)
	for c := 0; c < len(s); c++ {
		m[string(rune(s[c]))] = 1
	}

	r := []string{}
	for k := range m {
		r = append(r, k)
	}
	return r
}

func twoToOne(s1, s2 string) string {
	r1 := uniques(s1)
	r2 := uniques(s2)

	r3 := make([]string, len(r1)+len(r2))
	for _, v := range r1 {
		r3 = append(r3, v)
	}
	for _, v := range r2 {
		r3 = append(r3, v)
	}

	s3 := strings.Join(r3, "")

	r4 := uniques(s3)
	sort.Strings(r4)

	s4 := strings.Join(r4, "")
	return s4
}

func main() {
	s := twoToOne("asa", "esssesse")
	fmt.Println(s)
}
