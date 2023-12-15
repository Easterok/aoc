package main

import (
	"aoc/day15"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day15.Part1(content)

	fmt.Printf("Result: %d", result)
}
