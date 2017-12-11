package stream

import (
	"strings"
)

// CountScore return score and garbage count for given input
func CountScore(input string) (int, int) {
	data := strings.Split(input, "")
	isGarbage := false
	score := 0
	depth := 1
	garbageCount := 0
	var next string
	for i := 0; i < len(data); i++ {
		next = data[i]

		if next == "!" {
			i++
		} else if isGarbage && next != ">" {
			garbageCount++
		} else if next == "<" {
			isGarbage = true
		} else if next == ">" {
			isGarbage = false
		} else if next == "{" {
			score += depth
			depth++
		} else if next == "}" {
			depth--
		}
	}

	return score, garbageCount
}
