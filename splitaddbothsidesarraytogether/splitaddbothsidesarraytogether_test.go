package main

import (
	"reflect"
	"testing"
)

func TestSplitAndAdd(t *testing.T) {
	s := SplitAndAdd([]int{1, 2, 3, 4, 5}, 2)
	if !reflect.DeepEqual(s, []int{5, 10}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{5, 10})
	}

	s = SplitAndAdd([]int{15}, 3)
	if !reflect.DeepEqual(s, []int{15}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{15})
	}
	s = SplitAndAdd([]int{1, 2, 3, 4}, 1)
	if !reflect.DeepEqual(s, []int{4, 6}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{4, 6})
	}
	s = SplitAndAdd([]int{1, 2, 3, 4, 5, 6}, 20)
	if !reflect.DeepEqual(s, []int{21}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{21})
	}
	s = SplitAndAdd([]int{32, 45, 43, 23, 54, 23, 54, 34}, 2)
	if !reflect.DeepEqual(s, []int{183, 125}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{183, 125})
	}
	s = SplitAndAdd([]int{32, 45, 43, 23, 54, 23, 54, 34}, 0)
	if !reflect.DeepEqual(s, []int{32, 45, 43, 23, 54, 23, 54, 34}) {
		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{32, 45, 43, 23, 54, 23, 54, 34})
	}
	// 	s = SplitAndAdd([]int{3, 234, 25, 345, 45, 34, 234, 235, 345}, 3)
	// 	if !reflect.DeepEqual(s, []int{305, 1195}) {
	// 		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{305, 1195})
	// 	}
	// 	s = SplitAndAdd([]int{3, 234, 25, 345, 45, 34, 234, 235, 345, 34, 534, 45, 645, 645, 645, 4656, 45, 3}, 4)
	// 	if !reflect.DeepEqual(s, []int{1040, 7712}) {
	// 		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{1040, 7712})
	// 	}
	// 	s = SplitAndAdd([]int{23, 345, 345, 345, 34536, 567, 568, 6, 34536, 54, 7546, 456}, 20)
	// 	if !reflect.DeepEqual(s, []int{79327}) {
	// 		t.Errorf("SplitAndAdd was incorrect, got: %d, want %d.", s, []int{79327})
	// 	}
}
