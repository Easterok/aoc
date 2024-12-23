package aoc_2024

import (
	"fmt"
	"sort"
	"strings"
)

func Day23P1(content string) string {
	lines := strings.Split(content, "\n")
	connections := map[string]map[string]bool{}

	for _, line := range lines {
		a := line[0:2]
		b := line[3:]

		if _, ok := connections[a]; !ok {
			connections[a] = map[string]bool{}
		}
		connections[a][b] = true

		if _, ok := connections[b]; !ok {
			connections[b] = map[string]bool{}
		}
		connections[b][a] = true
	}

	groups := []string{}

	for a, mp := range connections {
		for b := range mp {
			for c := range connections[b] {
				if mp[c] && (strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t")) {
					acc := []string{a, b, c}
					sort.Slice(acc, func(i, j int) bool {
						return acc[i] < acc[j]
					})
					groups = append(groups, strings.Join(acc, "-"))
				}
			}
		}
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i] < groups[j]
	})

	c := map[string]bool{}

	for _, g := range groups {
		c[g] = true
	}

	return fmt.Sprintf("%d", len(c))
}

func day23search(curr string, network map[string]bool, connections map[string]map[string]bool) {
	if network[curr] {
		return
	}

	for key := range network {
		if !connections[key][curr] {
			return
		}
	}

	network[curr] = true

	for c := range connections[curr] {
		day23search(c, network, connections)
	}
}

func Day23P2(content string) string {
	lines := strings.Split(content, "\n")
	connections := map[string]map[string]bool{}

	for _, line := range lines {
		a := line[0:2]
		b := line[3:]

		if _, ok := connections[a]; !ok {
			connections[a] = map[string]bool{}
		}
		connections[a][b] = true

		if _, ok := connections[b]; !ok {
			connections[b] = map[string]bool{}
		}
		connections[b][a] = true
	}

	largest := map[string]bool{}

	for c := range connections {
		network := map[string]bool{}

		day23search(c, network, connections)

		if len(network) > len(largest) {
			largest = network
		}
	}

	keys := []string{}

	for k := range largest {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return strings.Join(keys, ",")
}
