package aoc_2024

import (
	"fmt"
	"strings"
)

func day19vartiants(design string, towels []string, cache *map[string]bool) bool {
	if len(design) == 0 {
		return true
	}

	if val, ok := (*cache)[design]; ok {
		return val
	}

	res := false

	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			in := day19vartiants(design[len(towel):], towels, cache)
			res = res || in
		}
	}

	(*cache)[design] = res

	return res
}

func Day19P1(content string) string {
	count := 0

	c := strings.Split(content, "\n\n")
	towels := strings.Split(c[0], ", ")
	designs := strings.Split(c[1], "\n")

	cache := map[string]bool{}

	for _, design := range designs {
		res := day19vartiants(design, towels, &cache)
		if res {
			count += 1
		}
	}

	return fmt.Sprintf("%d", count)
}

func day19vartiants2(design string, towels []string, cache *map[string]int) int {
	if len(design) == 0 {
		return 1
	}

	if val, ok := (*cache)[design]; ok {
		return val
	}

	res := 0

	for _, towel := range towels {
		if strings.HasPrefix(design, towel) {
			val := day19vartiants2(design[len(towel):], towels, cache)
			res += val
		}
	}

	(*cache)[design] = res

	return res
}

func Day19P2(content string) string {
	count := 0

	c := strings.Split(content, "\n\n")
	towels := strings.Split(c[0], ", ")
	designs := strings.Split(c[1], "\n")

	cache := map[string]int{}

	for _, design := range designs {
		val := day19vartiants2(design, towels, &cache)

		count += val
	}

	return fmt.Sprintf("%d", count)
}
