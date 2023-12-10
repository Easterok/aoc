package main

import (
	"aoc/day10"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day10.Part2(content)

	fmt.Printf("Result: %d", result)
}
