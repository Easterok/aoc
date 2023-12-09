package main

import (
	"aoc/day9"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day9.Part1(content)

	fmt.Printf("Result: %d", result)
}
