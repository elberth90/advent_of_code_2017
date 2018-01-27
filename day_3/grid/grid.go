package grid

import (
	"math"
)

// CalculateSteps calculate number of steps for given point using Manhattan Distance
func CalculateSteps(gridNo int) int {
	pos := calculateCoord(gridNo)
	steps := math.Abs(0-float64(pos[0])) + math.Abs(0-float64(pos[1]))
	return int(steps)
}

func calculateCoord(n int) [2]int {
	var pos [2]int

	k := int(math.Ceil((math.Sqrt(float64(n)) - 1) / 2))
	t := 2*k + 1
	m := int(math.Pow(float64(t), 2))

	t--

	if n >= m-t {
		pos[0] = k - (m - n)
		pos[1] = -k
		return pos
	}

	m -= t

	if n >= m-t {
		pos[0] = -k
		pos[1] = -k + (m - n)
		return pos
	}

	m -= t

	if n >= m-t {
		pos[0] = -k + (m - n)
		pos[1] = k
		return pos

	}

	pos[0] = k
	pos[1] = k - (m - n - t)

	return pos
}

const (
	right = iota
	up
	left
	down
)

// CalculateNext return first value that is larger than given input. Values are calculated using sum of the values in all adjacent squares, including diagonals.
// TODO: optimize!!
func CalculateNext(n int) int {
	g := make(map[int]map[int]int)
	g[0] = make(map[int]int)
	g[0][0] = 1

	c := 1
	b := 2
	dir := right
	x, y := 0, 0
	for i := 1; ; i++ {
		if c == b {
			switch dir {
			case right:
				dir = up
			case up:
				b++
				dir = left
			case left:
				dir = down
			case down:
				b++
				dir = right
			}
			c = 1
		}

		switch dir {
		case right:
			x++
		case up:
			y++
		case left:
			x--
		case down:
			y--
		}
		c++

		if g[x] == nil {
			g[x] = make(map[int]int)
		}
		g[x][y] = 0
		for ix := x - 1; ix <= x+1; ix++ {
			for iy := y - 1; iy <= y+1; iy++ {
				if ix == x && iy == y {
					continue
				}
				if _, ok := g[ix]; !ok {
					continue
				}

				if _, ok := g[ix][iy]; !ok {
					continue
				}
				g[x][y] += g[ix][iy]
			}
		}
		if g[x][y] > n {
			return g[x][y]
		}
	}
}

// normal spiral grid
//func spiralGrid(n int) {
//	c := 1
//	b := 2
//	dir := right
//	x, y := 0, 0
//	for i := 1; i <= n; i++ {
//		if c == b {
//			switch dir {
//			case right:
//				dir = up
//			case up:
//				b += 1
//				dir = left
//			case left:
//				dir = down
//			case down:
//				b += 1
//				dir = right
//			}
//			c = 1
//		}
//		switch dir {
//		case right:
//			x += 1
//		case up:
//			y += 1
//		case left:
//			x -= 1
//		case down:
//			y -= 1
//		}
//		c += 1
//	}
//}
