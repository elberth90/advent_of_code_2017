package main

import (
	"fmt"
	"io/ioutil"

	"strings"

	"github.com/elberth90/advent_of_code_2017/day_7/tower"
)

const filename = "data.txt"

func main() {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	result := tower.FindBottomProgram(strings.Trim(string(byteData), "\n"))
	fmt.Printf("Result : `%s`\n", result)

	result2 := tower.GetWeightToBalance(strings.Trim(string(byteData), "\n"))
	fmt.Printf("Result : `%d`\n", result2)
}
