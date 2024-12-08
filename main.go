package main

import (
	"aoc/2023/utils"
	aoc_2024 "aoc/2024"
	"fmt"
)

func main() {
	content := utils.ReadFile("./input")

	res := aoc_2024.Day8P2(content)

	fmt.Println("result:", res)
}
