package main

import (
	"aoc/day11"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day11.Part1(content)

	fmt.Printf("Result: %d", result)
}
