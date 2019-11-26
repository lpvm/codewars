package main

import "fmt"

// NbMonths is a function
func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	ret := [2]int{}
	months := 0
	addMonth := float64(savingperMonth)
	necessary := float64(startPriceNew)
	princeNew := float64(startPriceOld)
	savings := 0.0
	for {
		savings = float64(months) * addMonth
		if princeNew+savings >= necessary {
			ret[0] = months
			ret[1] = int(princeNew + savings - necessary)
			break
		}
		months++
		if months%2 == 0 && months != 0 {
			percentLossByMonth += 0.5
		}
		princeNew -= princeNew * percentLossByMonth / 100
		necessary -= necessary * percentLossByMonth / 100
	}
	return ret
}

func main() {
	nb := NbMonths(2000, 8000, 1000, 1.5)
	if nb == [2]int{6, 766} {
		fmt.Println(nb, "ok")
	} else {
		fmt.Println(nb, "NOT ok")
	}

}

// end month 1: percentLoss 1.5 available -4910.0
// end month 2: percentLoss 2.0 available -3791.7999...
// end month 3: percentLoss 2.0 available -2675.964
// end month 4: percentLoss 2.5 available -1534.06489...
// end month 5: percentLoss 2.5 available -395.71327...
// end month 6: percentLoss 3.0 available 766.158120825...
// return [6, 766] or (6, 766)
