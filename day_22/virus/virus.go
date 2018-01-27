package virus

import (
	"strings"
)

const (
	infected = "i"
	clean    = "c"
	weakened = "w"
	flagged  = "f"
	hash     = 35
	up       = 0
	right    = 1
	down     = 2
	left     = 3
)

type node struct {
	posX   int
	posY   int
	status string
}

// CalculateInfections return number of infections
func CalculateInfections(initialGrid string, bursts int) int {
	infectedList := loadInitiallyInfected(initialGrid)
	var face = up
	var posX = 0
	var posY = 0
	var infectedCounter = 0
	for i := 0; i < bursts; i++ {
		if index, isInfected := getIndexIfOnList(posX, posY, infectedList); isInfected {
			infectedList = append(infectedList[:index], infectedList[index+1:]...)
			face = (face + 4 + 1) % 4
		} else {
			infectedList = append(infectedList, node{posX: posX, posY: posY})
			infectedCounter++
			face = (face + 4 - 1) % 4
		}

		switch face {
		case up:
			posY++
		case right:
			posX++
		case down:
			posY--
		case left:
			posX--
		}
	}
	return infectedCounter
}

// CalculateInfectionsWithSpeedUp return number of infections after change of the logic
func CalculateInfectionsWithSpeedUp(initialGrid string, bursts int) int {
	nodeList := load(initialGrid)
	var infectedCounter = 0
	var status string
	var face = up
	var posX = 0
	var posY = 0
	for i := 0; i < bursts; i++ {
		if _, ok := nodeList[posX]; !ok {
			nodeList[posX] = map[int]*node{}
		}
		if _, ok := nodeList[posX][posY]; !ok {
			nodeList[posX][posY] = &node{posX: posX, posY: posY, status: clean}
		}

		n := nodeList[posX][posY]
		status = n.status

		switch n.status {
		case clean:
			n.status = weakened
		case weakened:
			n.status = infected
			infectedCounter++
		case infected:
			n.status = flagged
		case flagged:
			n.status = clean
		}

		switch status {
		case clean:
			face = (face + 4 - 1) % 4
		case infected:
			face = (face + 4 + 1) % 4
		case flagged:
			face = (face + 4 + 2) % 4
		}

		switch face {
		case up:
			posY++
		case right:
			posX++
		case down:
			posY--
		case left:
			posX--
		}
	}
	return infectedCounter
}

func getIndexIfOnList(posX int, posY int, infectedNodes []node) (int, bool) {
	var isInfected = false
	var index = 0
	for index, n := range infectedNodes {
		if n.posX != posX {
			continue
		}
		if n.posY != posY {
			continue
		}
		isInfected = true
		return index, isInfected
	}

	return index, isInfected
}

func loadInitiallyInfected(initialGrid string) []node {
	var infectedList []node
	var columns = len(strings.Split(initialGrid, "\n")[0])
	elements := strings.Replace(
		strings.Trim(initialGrid, "\n"),
		"\n",
		"",
		-1,
	)
	var rows = len(elements) / columns
	var startX = 0 - columns/2
	var startY = 0 + rows/2
	for _, element := range elements {
		if startX == columns/2+1 {
			startX = 0 - columns/2
			startY--
		}
		if element == hash {
			infectedList = append(infectedList, node{posX: startX, posY: startY})
		}
		startX++
	}

	return infectedList
}

func load(initialGrid string) map[int]map[int]*node {
	list := map[int]map[int]*node{}
	var columns = len(strings.Split(initialGrid, "\n")[0])
	elements := strings.Replace(
		strings.Trim(initialGrid, "\n"),
		"\n",
		"",
		-1,
	)
	var rows = len(elements) / columns
	var startX = 0 - columns/2
	var startY = 0 + rows/2
	for _, element := range elements {
		if startX == columns/2+1 {
			startX = 0 - columns/2
			startY--
		}

		if _, ok := list[startX]; !ok {
			list[startX] = map[int]*node{}
		}

		var n node
		if element == hash {
			n = node{posX: startX, posY: startY, status: infected}
		} else {
			n = node{posX: startX, posY: startY, status: clean}
		}

		list[startX][startY] = &n
		startX++
	}

	return list
}
