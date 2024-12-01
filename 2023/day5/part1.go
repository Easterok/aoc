package day5

import (
	"aoc/utils"
	"math"
	"strings"
)

type DSL struct {
	destination uint64
	source      uint64
	length      uint64
}

func lineToDSLMap(line_map string) (string, []DSL) {
	arrs := strings.Split(line_map, "\r\n")

	name := strings.ReplaceAll(arrs[0], " map:", "")

	dsl_result := []DSL{}

	for i := 1; i < len(arrs); i++ {
		raw_dsl := arrs[i]

		dsl_nums := utils.LineToUint64(raw_dsl)

		destination := dsl_nums[0]
		source := dsl_nums[1]
		length := dsl_nums[2]

		dsl_result = append(dsl_result, DSL{destination: destination, source: source, length: length})
	}

	return strings.TrimSpace(name), dsl_result
}

func findMapStartingWith(start string, dslmap map[string][]DSL) (string, []DSL) {
	for key, value := range dslmap {
		if strings.HasPrefix(key, start) {
			return key, value
		}
	}

	return "", []DSL{}
}

func getNextMapName(name string) string {
	return strings.Split(name, "-")[2]
}

func findSeedToLocation(start string, seed uint64, dslmap map[string][]DSL) uint {
	name, dsl := findMapStartingWith(start, dslmap)

	if start == "location" {
		return uint(seed)
	}

	current_name := name
	current_dsl := dsl

	next_seed := seed

	if len(current_dsl) > 0 && len(current_name) > 0 {
		for _, item := range current_dsl {
			source_end := item.source + item.length

			if item.source <= seed && seed < source_end {
				inc := seed - item.source

				next_seed = item.destination + inc
			}
		}
	}

	next_name := getNextMapName(name)

	return findSeedToLocation(next_name, next_seed, dslmap)
}

func Part1(content string) uint {
	lines := strings.Split(content, "\r\n\r\n")

	var result uint = math.MaxUint

	seedsLine := strings.Replace(strings.TrimSpace(lines[0]), "seeds: ", "", 1)

	seeds := utils.LineToUint64(seedsLine)

	dslmap := make(map[string][]DSL)

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		name, arrs := lineToDSLMap(line)

		dslmap[name] = arrs
	}

	for _, seed := range seeds {
		res := findSeedToLocation("seed", seed, dslmap)

		result = min(res, result)
	}

	return result
}
