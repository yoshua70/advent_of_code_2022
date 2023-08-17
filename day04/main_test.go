package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 2

	answer := PartOne(input)

	if answer != want {
		t.Fatalf(`PartOne("%s") = %d, want equal for %d`, input, answer, want)
	}
}

func TestPartTwo(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	want := 4

	answer := PartTwo(input)

	if answer != want {
		t.Fatalf(`PartTwo("%s") = %d, want equal for %d`, input, answer, want)
	}
}
