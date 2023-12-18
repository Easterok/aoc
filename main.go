package main

import (
	"aoc/day18"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day18.Part1(content)

	fmt.Printf("Result: %d", result)
}
