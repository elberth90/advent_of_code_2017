package path

import (
	"strings"
	"unicode"
)

type coord struct {
	x int
	y int
}

// Follow return letters that were left on the path and number of steps
func Follow(input string) (string, int) {
	diagram := strings.Split(strings.Trim(input, "\n"), "\n")
	directions := []coord{{x: 0, y: -1}, {x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: 1}}
	position := coord{x: findStart(diagram[0]), y: 0}
	direction := coord{x: 0, y: 1}
	path := ""
	count := 0
	for {
		var prevPosition coord
		prevPosition, position = position, coord{x: position.x + direction.x, y: position.y + direction.y}
		count++
		if position.x < 0 || position.x >= len(diagram[0]) || position.y < 0 || position.y >= len(diagram) {
			break
		}

		element := string(diagram[position.y][position.x])
		if isLetter(element) {
			path += element
		}

		if element == " " {
			break
		}

		if element == "+" {
			for _, possibleDirection := range directions {
				next := coord{x: position.x + possibleDirection.x, y: position.y + possibleDirection.y}
				if next == prevPosition {
					continue
				}

				if next.x < 0 || next.x >= len(diagram[0]) || next.y < 0 || next.y >= len(diagram) {
					continue
				}

				if string(diagram[next.y][next.x]) == " " {
					continue
				}
				direction = possibleDirection
			}
		}
	}
	return path, count
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func findStart(diagram string) int {
	for i, e := range diagram {
		if string(e) == "|" {
			return i
		}
	}
	panic("cannot find start")
}
