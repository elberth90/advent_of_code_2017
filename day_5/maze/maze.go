package maze

import (
	"strconv"
	"strings"
)

// FindExit return information about number of steps required to leave a maze
func FindExit(data string) (int, error) {
	jumpOffsets, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	index, steps, jump := 0, 0, 0
	for {
		jump = jumpOffsets[index]
		jumpOffsets[index]++
		index += jump
		steps++
		if index > len(jumpOffsets)-1 {
			break
		}
	}
	return steps, nil
}

func parseInput(data string) ([]int, error) {
	var err error
	d := strings.Split(strings.Trim(data, "\n"), "\n")
	jumpOffsets := make([]int, len(d))
	var i = 0
	for _, element := range d {
		jumpOffsets[i], err = strconv.Atoi(element)
		if err != nil {
			return nil, err
		}
		i++
	}

	return jumpOffsets, nil
}

// FindStrangeExit return information about number of steps required to leave a maze (strange jumps)
func FindStrangeExit(data string) (int, error) {

	jumpOffsets, err := parseInput(data)
	if err != nil {
		return 0, err
	}
	index, steps, jump := 0, 0, 0
	for {
		jump = jumpOffsets[index]
		if jump >= 3 {
			jumpOffsets[index]--
		} else {
			jumpOffsets[index]++
		}
		index += jump
		steps++
		if index > len(jumpOffsets)-1 {
			break
		}
	}
	return steps, nil
}
