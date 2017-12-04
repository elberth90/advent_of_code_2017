package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_2/checksum"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")

	sum, err := checksum.CalculateChecksum(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: `%d`\n", sum)

	sum, err = checksum.CalculateChecksumWithModulo(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result with modulo: `%d`\n", sum)
}
