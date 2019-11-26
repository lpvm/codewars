package main

import "fmt"

// GrowingPlant is a function
func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
	days := 0
	if upSpeed < 5 || upSpeed > 100 {
		return 0
	}
	if downSpeed < 2 || downSpeed > upSpeed {
		return 0
	}
	if desiredHeight < 4 || desiredHeight > 1000 {
		return 0
	}
	acum := upSpeed
	for {
		days++
		if acum >= desiredHeight {
			break
		}
		acum += upSpeed
		acum -= downSpeed
	}
	return days
}

func main() {
	gp0 := GrowingPlant(5, 2, 5)
	fmt.Println("gp0: ", gp0)
}

// Constraints: 5 ≤ upSpeed ≤ 100.
// Constraints: 2 ≤ downSpeed < upSpeed.
// Constraints: 4 ≤ desiredHeight ≤ 1000.
