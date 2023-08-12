package main

import (
	"fmt"
	"log"
	"yoshua70/helpers"
)

func main() {
	log.SetPrefix("calories_counting: ")
	log.SetFlags(0)

	input, err := helpers.GetInput("day2_input.txt")

	if err != nil {
		log.Fatal("Error opening input file.")
	}

	score := get_score(input)

	fmt.Println("Part one answer:", score)

}
