package main

import (
	"aoc/day12"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day12.Part2(content)

	fmt.Printf("Result: %d", result)
}
