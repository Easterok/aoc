package day5

import (
	"aoc/utils"
	"math"
	"strings"
)

type Range struct {
	start uint64
	end   uint64
}

func Part2(content string) uint {
	lines := strings.Split(content, "\r\n\r\n")

	var result uint = math.MaxUint

	seedsLine := strings.Replace(strings.TrimSpace(lines[0]), "seeds: ", "", 1)

	rawSeeds := utils.LineToUint64(seedsLine)

	seeds := []Range{}

	i := 0

	for i < len(rawSeeds) {
		start := rawSeeds[i]

		seeds = append(seeds, Range{start: start, end: start + rawSeeds[i+1]})
		i += 2
	}

	for _, line := range lines[1:] {
		splitted := strings.Split(line, "\r\n")[1:]
		ranges := [][]uint64{}

		for _, l := range splitted {
			ranges = append(ranges, utils.LineToUint64(l))
		}

		new_something := []Range{}

		for len(seeds) > 0 {
			seedRange := seeds[len(seeds)-1]
			seeds = seeds[:len(seeds)-1]

			wasBreak := false

			for _, r := range ranges {
				destination := r[0]
				source := r[1]
				length := r[2]

				overlapStart := max(seedRange.start, source)
				overlapEnd := min(seedRange.end, source+length)
				wasBreak = false

				if overlapStart < overlapEnd {
					new_something = append(new_something, Range{start: overlapStart - source + destination, end: overlapEnd - source + destination})

					if overlapStart > seedRange.start {
						seeds = append(seeds, Range{start: seedRange.start, end: overlapStart})
					}

					if overlapEnd < seedRange.end {
						seeds = append(seeds, Range{start: overlapEnd, end: seedRange.end})
					}

					wasBreak = true

					break
				}
			}

			if !wasBreak {
				new_something = append(new_something, seedRange)
			}
		}

		seeds = new_something
	}

	for _, seed := range seeds {
		result = min(uint(seed.start), result)
	}

	return result
}
