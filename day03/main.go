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

var ITEMS = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

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
		answer := partOne(input)
		fmt.Println("[Day 03] Answer:", answer)

	} else {
		answer := partTwo(input)
		fmt.Println("[Day 03] Answer:", answer)
	}
}

func partOne(input *os.File) int {
	prioritiesSum := 0

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		rucksack := strings.Split(scanner.Text(), "")
		misplacedItems := GetItemsInBothCompartment(SplitRucksack(rucksack))
		prioritiesSum += SumItemsPriorities(misplacedItems)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return prioritiesSum
}

func partTwo(input *os.File) int {
	count := 0
	result := 0
	scanner := bufio.NewScanner(input)
	// 	s1 := `vJrwpWtwJgWrhcsFMMfFFhFp
	// jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
	// PmmdzqPrVvPwwTWBwg`
	// 	s2 := `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
	// ttgJtRGJQctTZtZT
	// CrZsJsPPZsGzwwsLwLmpwMDw`

	// 	fmt.Println(getBadge(s1))
	// 	fmt.Println(getBadge(s2))

	s := ""
	for scanner.Scan() {
		s += scanner.Text() + "\n"
		if count == 2 {
			p, _ := ItemToPriority(getBadge(s))
			result += p
			s = ""
			count = 0
		} else {
			count++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func getBadge(rucksacks string) string {
	badge := ""
	rucksacks = strings.TrimRight(rucksacks, "\n")
	s := strings.Split(rucksacks, "\n")

	// Get the rucksack with the less items
	itemsCount := len(s[0])
	rucksackIndex := 0
	for i, v := range s {

		if len(v) < itemsCount {
			rucksackIndex = i
			itemsCount = len(v)
		}
	}

	// Get the two other rucksacks
	var newS []string
	for i, v := range s {
		if i != rucksackIndex {
			newS = append(newS, v)
		}
	}

	// Iterate over the items (characters) of the rucksack
	// with the less objects.
	smallestRucksack := s[rucksackIndex]
	for _, v := range smallestRucksack {
		if strings.Contains(newS[0], string(v)) && strings.Contains(newS[1], string(v)) {
			badge = string(v)
			break
		}
	}

	return badge
}

func SumItemsPriorities(items []string) int {
	prioritiesSum := 0

	for _, v := range items {
		itemPriority, _ := ItemToPriority(v)
		prioritiesSum += itemPriority
	}

	return prioritiesSum
}

func SplitRucksack(rucksack []string) ([]string, []string) {
	firstCompartment, secondCompartment, err := RucksackToCompartments(rucksack)

	if err != nil {
		fmt.Println(err.Error())
	}

	return firstCompartment, secondCompartment
}

func GetItemsInBothCompartment(firstCompartment, secondCompartment []string) []string {
	items := []string{}

	for i := 0; i < len(firstCompartment); i++ {
		for j := 0; j < len(secondCompartment); j++ {
			if firstCompartment[i] == secondCompartment[j] {
				if !arrayContains(items, firstCompartment[i]) {
					items = append(items, firstCompartment[i])
				}
			}
		}
	}

	return items
}

func arrayContains(array []string, item string) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}

	return false
}

func ItemToPriority(item string) (int, error) {

	for i, s := range ITEMS {
		if item == s {
			return i + 1, nil
		}
	}

	return -1, errors.New("unknown item")
}

func ItemsStringToArray(items string) ([]string, error) {

	if len(items) == 0 {
		return nil, errors.New("invalid string of length 0")
	}

	itemsArray := strings.Split(items, "")

	return itemsArray, nil
}

func RucksackToCompartments(rucksack []string) ([]string, []string, error) {
	rucksackSize := len(rucksack)

	if rucksackSize%2 != 0 {
		return nil, nil, errors.New("rucksack's size is odd")
	}

	firstCompartment := rucksack[0 : rucksackSize/2]
	secondCompartment := rucksack[rucksackSize/2:]

	return firstCompartment, secondCompartment, nil
}
