package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	helpers "github.com/yoshua70/aoc_helpers"
)

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or part 2")
	flag.Parse()
	fmt.Println("[Day 01] Calories Counting")
	fmt.Println("[Day 01] Running Part", part)

	input, err := helpers.GetInput("./resources/day01_input.txt")
	if err != nil {
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

func partOne(input *os.File) int {

	scanner := bufio.NewScanner(input)

	prev_calories, calories := 0, 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if prev_calories < calories {
				prev_calories = calories
			}
			calories = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			} else {
				calories += calorie
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return prev_calories
}

func partTwo(input *os.File) int {
	caloriesPerElf := getCaloriesPerElf(input)

	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})

	sumTopThreeCaloriesTotal := 0

	for i := 0; i <= 2; i++ {
		sumTopThreeCaloriesTotal += caloriesPerElf[i]
	}

	return sumTopThreeCaloriesTotal
}

func getCaloriesPerElf(input *os.File) []int {
	var caloriesPerElf []int
	var calories = 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		if scanner.Text() == "" {
			caloriesPerElf = append(caloriesPerElf, calories)
			calories = 0
		} else {
			calorie, err := strconv.Atoi(scanner.Text())
			if err == nil {
				calories += calorie
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

		return nil
	} else {
		return caloriesPerElf
	}
}
