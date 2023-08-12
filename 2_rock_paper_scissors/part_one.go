package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func get_score(input *os.File) int {
	scanner := bufio.NewScanner(input)

	score := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		score += get_shape_score(s[1])
		score += get_winning_score(s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}

// A, X -> Rock -> 1pt
// B, Y -> Paper -> 2pts
// C, Z -> Scissors -> 3pts
func get_shape_score(shape string) int {
	switch shape {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	return -1
}

// lost -> 0
// draw -> 3
// win -> 6
// In row, what the user plays.
// In column, what the elf plays.
//
//	C B A
//
// X 6 0 3
// Y 0 3 6
// Z 3 6 0
func get_winning_score(shapes []string) int {
	scores := [3][3]int{
		{6, 0, 3},
		{0, 3, 6},
		{3, 6, 0},
	}

	coor := get_matrix_coor(shapes)

	score := scores[coor[1]][coor[0]]

	return score
}

func get_matrix_coor(shapes []string) [2]int {
	var res [2]int

	if shapes[0] == "A" {
		res[0] = 2
	} else if shapes[0] == "B" {
		res[0] = 1
	} else if shapes[0] == "C" {
		res[0] = 0
	}

	if shapes[1] == "Z" {
		res[1] = 2
	} else if shapes[1] == "Y" {
		res[1] = 1
	} else if shapes[1] == "X" {
		res[1] = 0
	}

	return res
}
