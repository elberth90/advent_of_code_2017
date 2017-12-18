package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_18/parser"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := parser.FindFrequency(string(byteData))
	fmt.Printf("Result part1: `%d`\n", result)

	result = parser.Part2(string(byteData))
	fmt.Printf("Result part2: `%d`\n", result)

}
