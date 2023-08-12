package main

import (
	"bufio"
	"log"
	"strings"

	"github.com/yoshua70/advent_of_code_2022/helpers"
)

func get_score_2() int {
	input, err := helpers.GetInput("day2_input.txt")

	if err != nil {
		log.Fatal("Error opening input file.")
	}

	scanner := bufio.NewScanner(input)

	score := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		score += get_shape_points(what_to_play(s))
		score += get_play_points(s[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}

// A, X -> Rock -> 1pt
// B, Y -> Paper -> 2pts
// C, Z -> Scissors -> 3pts
func get_shape_points(shape string) int {
	switch shape {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	}

	return -1
}

func get_play_points(outcome string) int {
	switch outcome {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		return -1
	}
}

// X -> lose -> 0pt
// Y -> Draw -> 3pts
// Z -> Win -> 6pts
// A -> Rock -> 1pt
// B -> Paper -> 2pts
// C -> Scissors -> 3pts
func what_to_play(play []string) string {
	opponent_play := play[0]
	outcome_play := play[1]

	to_play := ""

	if outcome_play == "X" {
		to_play = lose_against(opponent_play)
	} else if outcome_play == "Y" {
		to_play = opponent_play
	} else {
		to_play = win_against(opponent_play)
	}

	return to_play
}

func win_against(play string) string {
	switch play {
	case "A":
		return "B"
	case "B":
		return "C"
	case "C":
		return "A"
	default:
		return ""
	}
}

func lose_against(play string) string {
	switch play {
	case "A":
		return "C"
	case "B":
		return "A"
	case "C":
		return "B"
	default:
		return ""
	}
}
