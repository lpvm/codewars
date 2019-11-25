package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringify(m map[string]int, l []string) string {
	s := "("
	for i, v := range l {
		if i != 0 {
			s += ") - ("
		}
		s += v + " : " + strconv.Itoa(m[v])
	}
	s = s + ")"
	return s
}

// StockList is a function
func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	m := map[string]int{}
	for _, c := range listCat {
		m[c] = 0
	}
	artCat := ""
	for _, l := range listArt {
		article := strings.Split(l, " ")
		artCat = string(rune(article[0][0]))
		_, ok := m[artCat]
		if ok {
			n, err := strconv.Atoi(article[1])
			if err == nil {
				m[artCat] += n
			}
		}
	}
	s := stringify(m, listCat)
	return s
}

func main() {
	b := []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	c := []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c))
	d := []string{"EBAR 0", "EDXE 0", "FKWR 0", "GTSQ 0", "XRTY 0"}
	e := []string{"A", "B", "C", "D"}
	fmt.Println(StockList(d, e))
}
