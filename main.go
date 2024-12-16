package main

import (
	"aoc/2023/utils"
	aoc_2024 "aoc/2024"
	"fmt"
	"os"
)

func main() {
	content := utils.ReadFile("./input")

	res := aoc_2024.Day15P2(content, len(os.Args) == 2 && os.Args[1] == "wasd")

	fmt.Println("result:", res)
}
