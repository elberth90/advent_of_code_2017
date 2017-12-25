package graph

import (
	"strconv"
	"strings"
)

// PartOne return result of part one
func PartOne(input string) int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	g := buildEdges(lines)
	programs := make(map[int]bool)

	var queue []int
	queue = append(queue, 0)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if _, ok := programs[v]; !ok {
			programs[v] = true
			queue = append(queue, g[v]...)
		}
	}

	return len(programs)
}

// PartTwo return result of part two
func PartTwo(input string) int {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	g := buildEdges(lines)
	programs := make(map[int]bool)

	groups := 0
	for k := range g {
		groups++
		var queue []int
		queue = append(queue, k)

		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			if _, ok := programs[v]; !ok {
				programs[v] = true
				queue = append(queue, g[v]...)
				delete(g, v)
			}
		}
	}

	return groups
}

func buildEdges(lines []string) map[int][]int {
	g := make(map[int][]int)
	for _, l := range lines {
		d := strings.Split(l, " <-> ")

		k, err := strconv.Atoi(d[0])
		if err != nil {
			panic(err)
		}
		values := strings.Split(d[1], ",")
		for _, v := range values {
			t, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				panic(err)
			}
			g[k] = append(g[k], t)
		}
	}
	return g
}
