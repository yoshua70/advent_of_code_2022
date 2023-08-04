package main

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"yoshua70/helpers"
)

/*
- input: the number of Calories each Elf is carrying
- output: the maximum number of Calories carried by an Elf
*/

func main() {
	log.SetPrefix("calories_counting: ")
	log.SetFlags(0)

	input, err := helpers.GetInput("day1_input.txt")

	if err != nil {
		log.Fatal("Error opening input file.")
	}

	scanner := bufio.NewScanner(input)

	var previous_calories_count, calories_count = 0, 0

	for scanner.Scan() {

		if scanner.Text() == "" {
			if previous_calories_count < calories_count {
				previous_calories_count = calories_count

			}
			calories_count = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			} else {
				calories_count += calorie
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(previous_calories_count)
}
