package main

import (
	"aoc/day7"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day7.Part1(content)

	fmt.Printf("Result: %d", result)
}
