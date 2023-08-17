package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "github.com/yoshua70/aoc_helpers"
)

func main() {
	input, err := helpers.GetInputAsString("../resources/day04_input.txt")

	if err != nil {
		fmt.Println(err)
		panic("Cannot open input file. Please check the file path.")
	}

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))
}

func PartOne(input string) int {
	res := 0

	pairs := strings.Split(input, "\n")

	for _, v := range pairs {
		sectionsRange := strings.Split(v, ",")

		sections := [][]int{
			mapStrToInt(strings.Split(sectionsRange[0], "-")),
			mapStrToInt(strings.Split(sectionsRange[1], "-")),
		}

		if sections[0][0] < sections[1][0] {
			if sections[0][1] >= sections[1][1] {
				res++
			}
		} else if sections[0][0] == sections[1][0] {
			res++
		} else if sections[0][0] > sections[1][0] {
			if sections[0][1] <= sections[1][1] {
				res++
			}
		}

	}

	return res
}

func PartTwo(input string) int {
	res := 0

	pairs := strings.Split(input, "\n")

	for _, v := range pairs {
		sectionsRange := strings.Split(v, ",")

		sections := [][]int{
			mapStrToInt(strings.Split(sectionsRange[0], "-")),
			mapStrToInt(strings.Split(sectionsRange[1], "-")),
		}

		if sections[0][1] >= sections[1][0] && sections[0][0] <= sections[1][1] {
			res++
		}
	}

	return res
}

func mapStrToInt(input []string) []int {
	var res []int

	for _, v := range input {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			panic("unable to convert string to integer")
		}
		res = append(res, intValue)
	}

	return res
}
