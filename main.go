package main

import (
	"aoc/day4"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day4.Part1(content)

	fmt.Printf("Result: %d", result)
}
