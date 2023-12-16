package main

import (
	"aoc/day1"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day1.Part2(content)

	fmt.Printf("Result: %d", result)
}
