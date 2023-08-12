package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("calories_counting: ")
	log.SetFlags(0)

	score := get_score()
	score_2 := get_score_2()

	fmt.Println("Part one answer:", score)
	fmt.Println("Part two answer:", score_2)
}
