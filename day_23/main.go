package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_23/coprocessor"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := coprocessor.HowManyMul(string(byteData))
	fmt.Printf("Result part1: `%d`\n", result)

	result = coprocessor.HRegisterValue()
	fmt.Printf("Result part2: `%d`\n", result)
}
