package main

import (
	"aoc/day15"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day15.Part2(content)

	fmt.Printf("Result: %d", result)
}
