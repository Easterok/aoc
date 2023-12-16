package main

import (
	"aoc/day16"
	"aoc/utils"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	result := day16.Part2(content)

	fmt.Printf("Result: %d", result)
}
