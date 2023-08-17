package main

import "testing"

func TestPartOne(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	want := "CMZ"

	res := PartOne(input)

	if res != want {
		t.Fatalf("Want `%s`, but got `%s`", want, res)
	}
}
