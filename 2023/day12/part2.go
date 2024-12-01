package day12

import (
	"aoc/utils"
	"fmt"
	"strings"
)

var cache = map[string]uint64{}

func toCacheKey(springs string, missed []uint64) string {
	key := springs

	for _, k := range missed {
		key += fmt.Sprintf("%c", k)
	}

	return key
}

func findCachedVariants(springs string, missed []uint64) uint64 {
	if springs == "" {
		if len(missed) == 0 {
			return 1
		}

		return 0
	}

	if len(missed) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}

		return 1
	}

	key := toCacheKey(springs, missed)

	cached, ok := cache[key]

	if ok {
		return cached
	}

	ch := springs[0]
	result := uint64(0)

	if ch == '.' || ch == '?' {
		result += findCachedVariants(springs[1:], missed)
	}

	if ch == '#' || ch == '?' {
		f := int(missed[0])

		if f <= len(springs) && !strings.Contains(springs[:f], ".") && (f == len(springs) || springs[f] != '#') {
			if f == len(springs) {
				result += findCachedVariants("", missed[1:])
			} else {
				result += findCachedVariants(springs[f+1:], missed[1:])
			}
		}
	}

	cache[key] = result

	return result
}

func Part2(content string) uint64 {
	lines := strings.Split(content, "\r\n")

	result := uint64(0)

	for _, line := range lines {
		splitted := strings.Split(line, " ")

		springs := splitted[0]

		missed := utils.LineToUint64(strings.ReplaceAll(splitted[1], ",", " "))

		s := []string{}
		nums := []uint64{}

		for i := 0; i < 5; i++ {
			s = append(s, springs)
			nums = append(nums, missed...)
		}

		result += findCachedVariants(strings.Join(s, "?"), nums)
	}

	return result
}
