package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_5/maze"
)

const filename = "data.txt"

func main() {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data := string(byteData)

	result, err := maze.FindExit(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %d\n", result)

	result, err = maze.FindStrangeExit(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %d\n", result)
}
