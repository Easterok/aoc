package aoc_2024

import (
	"fmt"
	"strings"
)

func day11evaluate(num, blinks int) int {
	if blinks == 0 {
		return 1
	}

	if num == 0 {
		return day11evaluate(1, blinks-1)
	}

	s := fmt.Sprintf("%d", num)

	if len(s)%2 == 0 {
		a := toInts(s[:len(s)/2])[0]
		b := toInts(s[len(s)/2:])[0]

		return day11evaluate(a, blinks-1) + day11evaluate(b, blinks-1)
	}

	return day11evaluate(num*2024, blinks-1)
}

func Day11P1(content string) string {
	count := 0

	blinks := 25
	nums := toInts(strings.Split(content, " ")...)

	for _, num := range nums {
		count += day11evaluate(num, blinks)
	}

	return fmt.Sprintf("%d", count)
}

type CachedBlink struct {
	value  int
	blinks int
}

func day11evaluate2(num, blinks int, cache *map[CachedBlink]int) int {
	cached := CachedBlink{value: num, blinks: blinks}

	if (*cache)[cached] != 0 {
		return (*cache)[cached]
	}

	if blinks == 0 {
		return 1
	}

	if num == 0 {
		(*cache)[cached] = day11evaluate2(1, blinks-1, cache)

		return (*cache)[cached]
	}

	s := fmt.Sprintf("%d", num)

	if len(s)%2 == 0 {
		aRes := day11evaluate2(toInts(s[:len(s)/2])[0], blinks-1, cache)
		bRes := day11evaluate2(toInts(s[len(s)/2:])[0], blinks-1, cache)

		(*cache)[cached] = aRes + bRes

		return (*cache)[cached]
	}

	(*cache)[cached] = day11evaluate2(num*2024, blinks-1, cache)

	return (*cache)[cached]
}

func Day11P2(content string) string {
	count := 0

	blinks := 75
	nums := toInts(strings.Split(content, " ")...)

	cache := map[CachedBlink]int{}

	for _, num := range nums {
		count += day11evaluate2(num, blinks, &cache)
	}

	return fmt.Sprintf("%d", count)
}
