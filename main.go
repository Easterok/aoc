package main

import (
	"aoc/day13"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day13.Part2(content)

	fmt.Printf("Result: %d", result)
}
