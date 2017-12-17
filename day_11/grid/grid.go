package grid

import (
	"math"
	"strings"
)

const (
//north = , northeast, southeast, south, southwest, and northwest
)

type coord struct {
	x int
	y int
	z int
}

func (c *coord) move(direction string) {
	switch direction {
	case "n":
		c.y++
		c.z--
	case "s":
		c.y--
		c.z++
	case "ne":
		c.x++
		c.z--
	case "sw":
		c.x--
		c.z++
	case "nw":
		c.y++
		c.x--
	case "se":
		c.y--
		c.x++
	}
}

// GetDistance return distance between origin at 0, 0, 0 and reached coord
func GetDistance(instruction string) int {
	lines := strings.Split(strings.Trim(instruction, "\n"), ",")
	c := new(coord)

	for _, line := range lines {
		c.move(line)
	}
	orig := new(coord)

	return distance(orig, c)
}

// GetLongestDistance return the furthest distance between origin and coord
func GetLongestDistance(instruction string) int {
	lines := strings.Split(strings.Trim(instruction, "\n"), ",")
	c := new(coord)
	orig := new(coord)
	var maxDistance int

	for _, line := range lines {
		c.move(line)
		distance := distance(orig, c)
		if distance > maxDistance {
			maxDistance = distance
		}
	}

	return maxDistance
}

func distance(c1, c2 *coord) int {
	return int(math.Abs(float64(c1.x-c2.x))+math.Abs(float64(c1.y-c2.y))+math.Abs(float64(c1.z-c2.z))) / 2
}
