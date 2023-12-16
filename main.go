package main

import (
	"aoc/day3"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day3.Part2(content)

	fmt.Printf("Result: %d", result)
}
