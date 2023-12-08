package day8

import (
	"fmt"
	"strings"
)

type Node struct {
	left  string
	right string
}

func makeDict(s []string) map[string]Node {
	res := map[string]Node{}

	for _, k := range s {
		tr := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(k), "(", ""), ")", "")
		node := strings.Split(tr, " = ")
		value := node[0]
		kv := strings.Split(node[1], ", ")

		res[value] = Node{left: kv[0], right: kv[1]}
	}

	return res
}

func Part1(content string) int {
	result := 0

	c := strings.Split(content, "\r\n")

	rawInstructions := strings.TrimSpace(c[0])
	instructions := strings.Split(rawInstructions, "")

	dict := makeDict(c[2:])

	current := "AAA"
	instructionIndex := 0

	for current != "ZZZ" {
		node := dict[current]
		goTo := instructions[instructionIndex]

		if goTo == "L" {
			current = node.left
		} else if goTo == "R" {
			current = node.right
		} else {
			fmt.Printf("Invalid instruction %s\n", goTo)
		}

		result += 1
		instructionIndex += 1

		if instructionIndex > len(instructions)-1 {
			instructionIndex = 0
		}
	}

	fmt.Println(dict)

	return result
}
