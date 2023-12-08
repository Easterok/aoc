package main

import (
	"aoc/day8"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day8.Part1(content)

	fmt.Printf("Result: %d", result)
}
