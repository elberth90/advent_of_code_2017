package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_1/captcha"
)

const filename = "data.txt"

func main() {

	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := string(byteData)
	data = strings.Trim(data, "\n")

	sum, err := captcha.CalculateSum(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result without overhead: `%d`\n", sum)

	sum, err = captcha.CalculateSumWithJump(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result with overhead: `%d`\n", sum)
}
