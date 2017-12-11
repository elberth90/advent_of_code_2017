package main

import (
	"io/ioutil"
	"strings"

	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_8/register"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")
	max, highestMax := register.Run(data)

	fmt.Printf("Result with overhead: `%d` `%d`\n", max, highestMax)
}
