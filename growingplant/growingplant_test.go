package main

import "testing"

func TestGrowingPlant(t *testing.T) {
	gp0 := GrowingPlant(5, 2, 5)
	r0 := 1
	if gp0 != r0 {
		t.Errorf("for (5,2,5) expected %d, got %d", 1, r0)
	}

	gp1 := GrowingPlant(100, 10, 910)
	r1 := 10
	if gp1 != r1 {
		t.Errorf("for (100,10,910) expected %d, got %d", 10, r1)
	}

}
