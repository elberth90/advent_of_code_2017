package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_10/hash"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := strings.Split(strings.Trim(string(byteData), "\n"), ",")
	inputLengths := make([]int, len(data))
	for index, i := range data {
		value, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputLengths[index] = value
	}

	result := hash.Part1(256, inputLengths)
	fmt.Printf("Result part1: `%d`\n", result)

	byteData = []byte(strings.Trim(string(byteData), "\n"))
	h := hash.Part2(256, byteData)
	fmt.Printf("Result part2: `%s`\n", h)
}
