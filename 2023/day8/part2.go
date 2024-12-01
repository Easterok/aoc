package day8

import (
	"fmt"
	"strings"
)

func makeStartingNodes(c map[string]Node) []string {
	res := []string{}

	for key := range c {
		if strings.HasSuffix(key, "A") {
			res = append(res, key)
		}
	}

	return res
}

func has(acc []string, item string) bool {
	if len(acc) == 0 {
		return false
	}

	for _, k := range acc {
		if k == item {
			return false
		}
	}

	return true
}

func gdc(a uint64, b uint64) uint64 {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a uint64, b uint64) uint64 {
	return a * b / gdc(a, b)
}

func lcm_list(a []uint64) uint64 {
	r := a[0]

	for _, k := range a[1:] {
		r = lcm(r, k)
	}

	return r
}

func Part2(content string) uint64 {
	c := strings.Split(content, "\r\n")

	rawInstructions := strings.TrimSpace(c[0])
	instructions := strings.Split(rawInstructions, "")

	dict := makeDict(c[2:])

	current := makeStartingNodes(dict)

	a := []uint64{}

	for _, item := range current {
		res := uint(0)
		instructionIndex := 0

		curr := item

		acc := []string{}

		for !has(acc, curr) {
			if strings.HasSuffix(curr, "Z") {
				break
			}

			goTo := instructions[instructionIndex]
			node := dict[curr]

			if goTo == "L" {
				curr = node.left
			} else if goTo == "R" {
				curr = node.right
			}

			instructionIndex += 1
			res += 1

			if instructionIndex > len(instructions)-1 {
				instructionIndex = 0
			}
		}

		a = append(a, uint64(res))
	}

	result := uint64(1)

	fmt.Println(a)

	res := lcm_list(a)

	fmt.Println(res)

	return result
}
