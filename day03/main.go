package main

import (
	"flag"
	"fmt"

	helpers "github.com/yoshua70/aoc_helpers"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()
	fmt.Println("[Day 03] Rucksack Reorganization")
	fmt.Println("[Day 03] Running Part", part)

	input, err := helpers.GetInput("../resources/day03_input.txt")
	if err != nil {
		fmt.Println(err.Error())
		panic("Error while opening input file.")
	}
	defer input.Close()

	if part == 1 {
		answer := ""
		fmt.Println("[Day 03] Answer:", answer)

	} else {
		answer := ""
		fmt.Println("[Day 03] Answer:", answer)
	}
}
