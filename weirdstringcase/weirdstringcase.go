package main

import (
	"fmt"
	"strings"
)

func toWeirdCase(str string) string {
	weird := []string{}
	orig := []rune{}

	s := strings.Split(str, " ")

	for _, k := range s {
		orig = []rune(k)
		w := []string{}
		for i, l := range orig {
			if i%2 == 0 {
				w = append(w, strings.ToUpper(string(l)))
			} else {
				w = append(w, strings.ToLower(string(l)))
			}
		}
		weird = append(weird, string(strings.Join(w, "")))
	}

	return strings.Join(weird, " ")
}

func main() {
	w := toWeirdCase("Weird string case")
	fmt.Println("w: ", w)
	y := toWeirdCase("ThIs iS A TeSt lOoKs lIkE YoU PaSsEd")
	//	                  ThIs iS A TeSt lOoKs lIkE YoU PaSsEd
	fmt.Println("y: ", y)
	fmt.Println("y:  ThIs Is A TeSt LoOkS LiKe YoU PaSsEd" == y)
}

// ThIs iS A TeSt lOoKs lIkE YoU PaSsEd
// ThIs Is A TeSt LoOkS LiKe YoU PaSsEd")
//toWeirdCase("Weird string case") // => returns "WeIrD StRiNg CaSe"
