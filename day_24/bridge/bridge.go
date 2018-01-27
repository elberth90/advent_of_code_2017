package bridge

import (
	"strconv"
	"strings"
)

type element struct {
	p1   int
	p2   int
	used bool
}

// StrongestBridge return the strength of the strongest bridge
func StrongestBridge(input string) int {
	elementList := parseInput(input)
	maxSum := 0
	buildForStrength(elementList, 0, 0, &maxSum)
	return maxSum
}

// LongestBridge return the strength of the longest bridge
func LongestBridge(input string) int {
	elementList := parseInput(input)
	maxLen := 0
	maxSum := 0
	buildForLength(elementList, 0, 0, &maxSum, 0, &maxLen)

	return maxSum
}

func buildForStrength(elements []element, port int, sum int, maxSum *int) {
	if sum > *maxSum {

		*maxSum = sum
	}
	for i, e := range elements {
		if e.used {
			continue
		}
		var cp = make([]element, len(elements))
		copy(cp, elements)
		cp[i].used = true
		if port == e.p1 {
			buildForStrength(cp, e.p2, sum+e.p1+e.p2, maxSum)
		} else if port == e.p2 {
			buildForStrength(cp, e.p1, sum+e.p1+e.p2, maxSum)
		}
	}
}

func buildForLength(elements []element, port int, sum int, maxSum *int, length int, maxLen *int) {
	if length >= *maxLen {
		*maxLen = length
		if sum > *maxSum {
			*maxSum = sum
		}
	}
	for i, e := range elements {
		if e.used {
			continue
		}
		var cp = make([]element, len(elements))
		copy(cp, elements)
		cp[i].used = true
		if port == e.p1 {
			buildForLength(cp, e.p2, sum+e.p1+e.p2, maxSum, length+1, maxLen)
		} else if port == e.p2 {
			buildForLength(cp, e.p1, sum+e.p1+e.p2, maxSum, length+1, maxLen)
		}
	}
}

func parseInput(input string) []element {
	data := strings.Split(strings.Trim(input, "\n"), "\n")
	var elementList = make([]element, len(data))
	for i, row := range data {
		e := strings.Split(row, "/")
		p1, err := strconv.Atoi(e[0])
		if err != nil {
			panic(err)
		}
		p2, err := strconv.Atoi(e[1])
		if err != nil {
			panic(err)
		}
		elementList[i] = element{p1: p1, p2: p2}
	}

	return elementList
}
