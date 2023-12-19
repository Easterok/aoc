package main

import (
	"aoc/day19"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day19.Part1(content)

	fmt.Printf("Result: %d", result)
}
