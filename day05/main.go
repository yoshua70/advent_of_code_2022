package main

import (
	"fmt"

	helpers "github.com/yoshua70/aoc_helpers"
)

func main() {

	input, err := helpers.GetInputAsString("../resources/day05_input.txt")

	if err != nil {
		panic("cannot open input file, please check file path")
	}

	fmt.Println(PartOne(input))
}

func PartOne(input string) string {
	res := ""

	return res
}
