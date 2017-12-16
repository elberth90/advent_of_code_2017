package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_16/dance"
)

const filename = "data.txt"

const times = 1000000000

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")
	result := dance.OneTime(16, data)
	fmt.Printf("Result: `%s`\n", result)

	result = dance.NTimes(16, data, times)
	fmt.Printf("Result: `%s`\n", result)

}
