package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/yoshua70/advent_of_code_2022/helpers"
)

func total_top_three_calories() int {
	log.SetPrefix("calories_counting: ")
	log.SetFlags(0)

	input, err := helpers.GetInput("day1_input.txt")

	if err != nil {
		log.Fatal("Error opening input file.")

		return -1
	} else {
		var calories_sums_array = calories_sums_per_elf(input)

		sort.Slice(calories_sums_array, func(i, j int) bool {
			return calories_sums_array[i] > calories_sums_array[j]
		})

		top_three_calories_sum := 0
		for i := 0; i <= 2; i++ {
			top_three_calories_sum += calories_sums_array[i]
		}

		return top_three_calories_sum

	}
}

// Take a pointer to an input file.
// Returns an array of integers corresponding to the sum of the calories
// carried by each elf.
func calories_sums_per_elf(input *os.File) []int {
	var calories_sums_array []int

	var calories_count = 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {

		if scanner.Text() == "" {
			calories_sums_array = append(calories_sums_array, calories_count)
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

		return nil
	} else {
		return calories_sums_array
	}

}
