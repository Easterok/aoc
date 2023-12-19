package day19

import (
	"aoc/utils"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Rule struct {
}

func parseRules(content string) map[string][]string {
	acc := map[string][]string{}

	splitted := strings.Split(content, "\r\n")

	for _, k := range splitted {
		c := strings.Split(k, "{")
		key := c[0]
		rules := strings.Split(strings.ReplaceAll(c[1], "}", ""), ",")

		acc[key] = rules
	}

	return acc
}

func extractNum(c string) int {
	return utils.LineToNums(strings.Split(c, "=")[1])[0]
}

func parseParts(content string) []map[string]int {
	lines := strings.Split(content, "\r\n")
	acc := []map[string]int{}

	for _, line := range lines {
		prep := strings.ReplaceAll(strings.ReplaceAll(line, "}", ""), "{", "")
		arr := strings.Split(prep, ",")

		x, m, a, s := extractNum(arr[0]), extractNum(arr[1]), extractNum(arr[2]), extractNum(arr[3])

		acc = append(acc, map[string]int{"x": x, "m": m, "a": a, "s": s})
	}

	return acc
}

func isAccepted(prt map[string]int, m map[string][]string, key string) bool {
	rules := m[key]

	for _, r := range rules {
		if strings.Contains(r, ":") {
			ss := strings.Split(r, ":")
			condition := ss[0]
			direction := ss[1]

			if strings.Contains(condition, ">") {
				more := strings.Split(condition, ">")

				val := prt[more[0]]

				if val > utils.LineToNums(more[1])[0] {
					if direction == "A" {
						return true
					}

					if direction == "R" {
						return false
					}

					return isAccepted(prt, m, direction)
				}
			}

			if strings.Contains(condition, "<") {
				more := strings.Split(condition, "<")

				val := prt[more[0]]

				if val < utils.LineToNums(more[1])[0] {
					if direction == "A" {
						return true
					}

					if direction == "R" {
						return false
					}

					return isAccepted(prt, m, direction)
				}
			}

			continue
		}

		if r == "A" {
			return true
		}

		if r == "R" {
			return false
		}

		return isAccepted(prt, m, r)
	}

	return false
}

func Part1(content string) int {
	s := strings.Split(content, "\r\n\r\n")

	rules, parts := s[0], s[1]

	parsedRules := parseRules(rules)
	parsedParts := parseParts(parts)

	total := 0

	for _, part := range parsedParts {
		if isAccepted(part, parsedRules, "in") {
			total += (part["x"] + part["m"] + part["a"] + part["s"])
		}
	}

	return total
}
