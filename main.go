package main

import (
	"aoc/day2"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input_test")

	result := day2.Part1(content)

	fmt.Printf("Result: %d", result)
}
