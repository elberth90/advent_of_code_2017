package main

import (
	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_14/disk"
)

const hashInput = "uugsqrei"

func main() {

	result := disk.Occupancy(hashInput)
	fmt.Printf("Result part1: `%d`\n", result)

	regions := disk.RegionsFinder(hashInput)
	fmt.Printf("Result part2: `%d`\n", regions)
}
