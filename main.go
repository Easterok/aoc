package main

import (
	"aoc/day17"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day17.Part1(content)

	fmt.Printf("Result: %d", result)
}
