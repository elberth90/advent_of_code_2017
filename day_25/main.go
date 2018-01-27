package main

import (
	"io/ioutil"

	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_25/parser"
	"github.com/elberth90/advent_of_code_2017/day_25/stateMachine"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	blueprint := parser.ParseBlueprint(string(byteData))
	result := stateMachine.Diagnostic(blueprint)
	fmt.Printf("Result part1: `%d`\n", result)

}
