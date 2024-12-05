package aoc_2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day5valid(s []string, rules map[string]*Rule) string {
	cp := make([]string, len(s))
	copy(cp, s)

	sort.Slice(cp, func(i, j int) bool {
		si := string(cp[i])
		sj := string(cp[j])

		if rules[si].Before[sj] || rules[sj].After[si] {
			return false
		}

		return true
	})

	if strings.Join(cp, ",") == strings.Join(s, ",") {
		return s[len(s)/2]
	}

	return "0"
}

func day5valid2(s []string, rules map[string]*Rule) string {
	cp := make([]string, len(s))
	copy(cp, s)

	sort.Slice(cp, func(i, j int) bool {
		si := string(cp[i])
		sj := string(cp[j])

		if rules[si].Before[sj] || rules[sj].After[si] {
			return false
		}

		return true
	})

	if strings.Join(cp, ",") != strings.Join(s, ",") {
		return cp[len(cp)/2]
	}

	return "0"
}

type Rule struct {
	Before map[string]bool
	After  map[string]bool
}

func day5rules(s string) map[string]*Rule {
	res := map[string]*Rule{}
	spl := strings.Split(s, "\n")

	for _, line := range spl {
		l := strings.Split(line, "|")

		item, ok := res[l[0]]

		if !ok {
			res[l[0]] = &Rule{
				After: map[string]bool{
					l[1]: true,
				},
				Before: map[string]bool{},
			}
		} else {
			item.After[l[1]] = true
		}

		item, ok = res[l[1]]

		if !ok {
			res[l[1]] = &Rule{
				After: map[string]bool{},
				Before: map[string]bool{
					l[0]: true,
				},
			}
		} else {
			item.Before[l[0]] = true
		}
	}

	return res
}

func Day5P1(content string) string {
	count := 0

	a := strings.Split(content, "\n\n")
	rulesContent := strings.TrimSpace(a[0])
	pageContent := strings.TrimSpace(a[1])

	rules := day5rules(rulesContent)
	pages := strings.Split(pageContent, "\n")

	for _, line := range pages {
		l := strings.Split(line, ",")

		res := day5valid(l, rules)
		num, _ := strconv.Atoi(res)
		count += num
	}

	return fmt.Sprintf("%d", count)
}

func Day5P2(content string) string {
	count := 0

	a := strings.Split(content, "\n\n")
	rulesContent := strings.TrimSpace(a[0])
	pageContent := strings.TrimSpace(a[1])

	rules := day5rules(rulesContent)
	pages := strings.Split(pageContent, "\n")

	for _, line := range pages {
		l := strings.Split(line, ",")

		res := day5valid2(l, rules)
		num, _ := strconv.Atoi(res)
		count += num
	}

	return fmt.Sprintf("%d", count)
}
