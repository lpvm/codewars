package main

import (
	"testing"
)

func TestToWeirdCase(t *testing.T) {

	t0 := toWeirdCase("abc def")
	e0 := "AbC DeF"
	if t0 != e0 {
		t.Errorf("for %s, expected: %s, got %s", "abc def", "AbC DeF", e0)
	}

	t0 = toWeirdCase("ABC")
	e0 = "AbC"
	if t0 != e0 {
		t.Errorf("for %s, expected: %s, got %s", "ABC", "AbC", e0)
	}

	t0 = toWeirdCase("This is a test Looks like you passed")
	e0 = "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd"
	if t0 != e0 {
		t.Errorf("for %s, expected: %s, got %s", "This is a test Looks like you passed", "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd", e0)
	}

}
