package aoc_2024

import (
	"fmt"
	"sort"
	"strings"
)

var day24op = map[string]func(a, b int) int{
	"AND": func(a, b int) int {
		return a & b
	},
	"OR": func(a, b int) int {
		return a | b
	},
	"XOR": func(a, b int) int {
		return a ^ b
	},
}

type Day24Comp struct {
	a      string
	b      string
	op     func(a, b int) int
	cached int
}

func day24solve(key string, acc map[string]Day24Comp) int {
	c := acc[key]

	if c.cached != -1 {
		return c.cached
	}

	a := acc[c.a]
	b := acc[c.b]

	if a.cached == -1 {
		a.cached = day24solve(c.a, acc)
	}

	if b.cached == -1 {
		b.cached = day24solve(c.b, acc)
	}

	val := c.op(a.cached, b.cached)
	c.cached = val

	return val
}

func Day24P1(content string) string {
	lines := strings.Split(content, "\n\n")

	init := strings.Split(lines[0], "\n")
	computed := strings.Split(lines[1], "\n")

	acc := map[string]Day24Comp{}

	shouldBeSolved := []string{}

	for _, comp := range computed {
		spl := strings.Split(comp, " ")
		a, op, b, c := spl[0], spl[1], spl[2], spl[4]
		opFn := day24op[op]

		acc[c] = Day24Comp{
			a:      a,
			b:      b,
			op:     opFn,
			cached: -1,
		}

		if strings.HasPrefix(c, "z") {
			shouldBeSolved = append(shouldBeSolved, c)
		}
	}

	for _, i := range init {
		a := strings.Split(i, ": ")
		cached := 0
		if a[1] == "1" {
			cached = 1
		}

		acc[a[0]] = Day24Comp{
			cached: cached,
		}
	}

	sort.Slice(shouldBeSolved, func(i, j int) bool {
		return shouldBeSolved[i] > shouldBeSolved[j]
	})

	res := int(0)

	for index, key := range shouldBeSolved {
		val := day24solve(key, acc)

		res |= val << (len(shouldBeSolved) - index - 1)
	}

	return fmt.Sprintf("%d", res)
}
