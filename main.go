package main

import (
	"aoc/day5"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day5.Part2(content)

	fmt.Printf("Result: %d", result)
}
