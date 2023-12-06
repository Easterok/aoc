package main

import (
	"aoc/day6"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day6.Part1(content)

	fmt.Printf("Result: %d", result)
}
