package main

import (
	"aoc/day2"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day2.Part1(content)

	fmt.Printf("Result: %d", result)
}
