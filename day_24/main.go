package main

import (
	"fmt"
	"io/ioutil"

	"github.com/elberth90/advent_of_code_2017/day_24/bridge"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := bridge.StrongestBridge(string(byteData))
	fmt.Printf("Result part1: `%d`\n", result)

	result = bridge.LongestBridge(string(byteData))
	fmt.Printf("Result part2: `%d`\n", result)

}
