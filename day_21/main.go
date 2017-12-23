package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_21/image"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := image.MakeArt(string(byteData), 5)
	fmt.Printf("Result part1: `%d`\n", result)

	result = image.MakeArt(string(byteData), 18)
	fmt.Printf("Result part2: `%d`\n", result)
}
