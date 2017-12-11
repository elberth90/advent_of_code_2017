package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_9/stream"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")
	score, garbageCount := stream.CountScore(data)
	fmt.Printf("Result: `%d` `%d`\n", score, garbageCount)

}
