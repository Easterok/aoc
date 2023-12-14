package main

import (
	"aoc/day14"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day14.Part1(content)

	fmt.Printf("Result: %d", result)
}
