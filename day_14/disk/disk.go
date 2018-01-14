package disk

import (
	"fmt"
	"strconv"
	"strings"
)

const gridLength = 128

// Occupancy return number of occupied grid squares
func Occupancy(input string) int {
	var occupancy int

	for i := 0; i < gridLength; i++ {
		hash := fmt.Sprintf("%s-%d", input, i)
		knotHash := Part2(256, []byte(hash))
		occupancy += strings.Count(convertToBits(knotHash), "1")
	}

	return occupancy
}

// RegionsFinder return number of regions in disk grid
func RegionsFinder(input string) int {
	var disk [gridLength]string
	var visited []bool = make([]bool, gridLength*gridLength)
	var maxGroup int

	for i := 0; i < gridLength; i++ {
		hash := fmt.Sprintf("%s-%d", input, i)
		knotHash := Part2(256, []byte(hash))
		disk[i] = convertToBits(knotHash)
	}

	for i := 0; i < gridLength; i++ {
		for j := 0; j < gridLength; j++ {
			if string(disk[i][j]) == "1" && !visited[i*gridLength+j] {
				visit(i, j, visited, disk)
				maxGroup++
			}
		}
	}

	return maxGroup
}

func visit(i int, j int, visited []bool, disk [gridLength]string) {
	if i < 0 || i >= gridLength || j < 0 || j >= 128 || visited[i*gridLength+j] || string(disk[i][j]) == "0" {
		return
	}

	visited[i*gridLength+j] = true
	visit(i-1, j, visited, disk)
	visit(i+1, j, visited, disk)
	visit(i, j-1, visited, disk)
	visit(i, j+1, visited, disk)
}

func convertToBits(knotHash string) string {
	var bits string
	for _, element := range knotHash {
		b, err := hexToBin(string(element))
		if err != nil {
			panic(err)
		}
		bits = fmt.Sprintf("%s%s", bits, b)
	}
	return bits
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04b", ui), nil
}
