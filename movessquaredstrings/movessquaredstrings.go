package main

import (
	"fmt"
	"strings"
)

func reverseHor(s string) string {
	fmt.Println("reverseHor s: ", s)
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(rune(s[i]))
	}
	return r
}

// VertMirror should perform the transformation
// "abcd\nefgh\nijkl\nmnop" -->  "dcba\nhgfe\nlkji\nponm"
func VertMirror(s string) string {
	sl := strings.Split(s, "\n")
	rev := []string{}
	for _, v := range sl {
		rev = append(rev, reverseHor(v))
	}
	r := strings.Join(rev, "\n")
	return r
}

// HorMirror should perform the transformation
// "abcd\nefgh\nijkl\nmnop" -->  "mnop\nijkl\nefgh\nabcd"
func HorMirror(s string) string {
	sl := strings.Split(s, "\n")
	rev := []string{}
	for i := len(sl) - 1; i >= 0; i-- {
		rev = append(rev, sl[i])
	}

	return strings.Join(rev, "\n")
}

// FParam is a type of a function that transforms a
// string into another string
type FParam func(string) string

// Oper calls the function HorMirror or VertMirror
func Oper(f FParam, x string) string {
	r := f(x)
	return r
}

func main() {
	s := "abcd\nefgh\nijkl\nmnop"
	vert_mirror := func(s string) string { return VertMirror(s) }
	hor_mirror := func(s string) string { return HorMirror(s) }
}

//s = "abcd\nefgh\nijkl\nmnop"
//oper(vert_mirror, s) => "dcba\nhgfe\nlkji\nponm"
//oper(hor_mirror, s) => "mnop\nijkl\nefgh\nabcd"
