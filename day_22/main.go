package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_22/virus"
)

const filename = "data.txt"
const bursts = 10000
const burstSpeedUp = 10000000

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := virus.CalculateInfections(string(byteData), bursts)
	fmt.Printf("Result part1: `%d`\n", result)

	result = virus.CalculateInfectionsWithSpeedUp(string(byteData), burstSpeedUp)
	fmt.Printf("Result part2: `%d`\n", result)
}
