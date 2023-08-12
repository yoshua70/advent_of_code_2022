package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	helpers "github.com/yoshua70/aoc_helpers"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()
	fmt.Println("[Day 02] Rock Paper Scissors")
	fmt.Println("[Day 02] Running Part", part)

	input, err := helpers.GetInput("../resources/day02_input.txt")
	if err != nil {
		fmt.Println(err.Error())
		panic("Error while opening input file.")
	}
	defer input.Close()

	if part == 1 {
		answer := partOne(input)
		fmt.Println("[Day 01] Answer:", answer)

	} else {
		answer := partTwo(input)
		fmt.Println("[Day 01] Answer:", answer)
	}
}

/*
A, X -> Rock -> 1pt
B, Y -> Paper -> 2pts
C, Z -> Scissors -> 3pts

lost -> 0pt
draw -> 3
win -> 6
*/
func partOne(input *os.File) int {
	score := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		score += GetPlayScore(s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}

/*
A, X -> Rock -> 1pt
B, Y -> Paper -> 2pts
C, Z -> Scissors -> 3pts

X -> lost -> 0pt
Y -> draw -> 3
Z -> win -> 6
*/
func partTwo(input *os.File) int {
	score := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		score += PlayToScore(s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}

func PlayToScore(play []string) int {
	outcomeScore, _ := OutcomeToScore(play[1])
	userShape := ShapeUserShouldPlay(play)

	userShapeScore, _ := ShapeToScore(userShape)

	return outcomeScore + userShapeScore
}

func ShapeToWinAgainst(shape string) string {
	switch shape {
	case "rock":
		return "paper"
	case "paper":
		return "scissors"
	case "scissors":
		return "rock"
	}

	return ""
}

func ShapeToLoseAgainst(shape string) string {
	switch shape {
	case "rock":
		return "scissors"
	case "paper":
		return "rock"
	case "scissors":
		return "paper"
	}

	return ""
}

func ShapeUserShouldPlay(play []string) string {
	outcome := play[1]
	opponentShape, _ := ShapeToString(play[0])
	res := ""

	if outcome == "X" {
		res = ShapeToLoseAgainst(opponentShape)
	} else if outcome == "Y" {
		res = opponentShape
	} else {
		res = ShapeToWinAgainst(opponentShape)
	}

	return res
}

func OutcomeToScore(outcome string) (int, error) {
	score := -1
	switch outcome {
	case "X":
		score = 0
	case "Y":
		score = 3
	case "Z":
		score = 6
	}

	if score == -1 {
		return -1, errors.New("unknown outcome")
	}

	return score, nil
}

/*
A matrix holds the outcome of the play,
In column, the shape played by the elf,
In row, the shape the user plays.

		A B C
	 X  3 0 6
	 Y	6 3 0
	 Z	0 6 3

To get the score, we get the coordinates corresponding in the matrix.
*/
func GetPlayScore(shapes []string) int {
	scores := [3][3]int{
		{3, 0, 6},
		{6, 3, 0},
		{0, 6, 3},
	}

	elfShape, _ := ShapeToString(shapes[0])
	userShape, _ := ShapeToString(shapes[1])

	col, _ := ShapeToCoor(elfShape)
	row, _ := ShapeToCoor(userShape)

	userShapeScore, _ := ShapeToScore(userShape)

	return scores[row][col] + userShapeScore
}

func ShapeToString(shape string) (string, error) {
	shapeString := ""

	switch shape {
	case "A", "X":
		shapeString = "rock"
	case "B", "Y":
		shapeString = "paper"
	case "C", "Z":
		shapeString = "scissors"
	}

	if shapeString == "" {
		return "", errors.New("unknown shape")
	}

	return shapeString, nil
}

func ShapeToCoor(shape string) (int, error) {
	coor := -1

	switch shape {
	case "rock":
		coor = 0
	case "paper":
		coor = 1
	case "scissors":
		coor = 2
	}

	if coor == -1 {
		return -1, errors.New("unknown shape")
	}

	return coor, nil
}

func ShapeToScore(shape string) (int, error) {
	score := -1

	switch shape {
	case "rock":
		score = 1
	case "paper":
		score = 2
	case "scissors":
		score = 3
	}

	if score == -1 {
		return -1, errors.New("unknown shape")
	}

	return score, nil
}
