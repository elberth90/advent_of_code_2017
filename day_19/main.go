package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_19/path"
)

const filename = "data.txt"

func main() {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	p, steps := path.Follow(string(byteData))
	fmt.Printf("Result part1: `%s`\n", p)
	fmt.Printf("Result part2: `%d`\n", steps)
}
