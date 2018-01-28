package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/elberth90/advent_of_code_2017/day_6/bank"
)

const filename = "data.txt"

func main() {
	byteData, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	inputList := strings.Fields(string(byteData))
	bankValues := make([]int, len(inputList))
	for i, value := range inputList {
		bankValues[i], err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	}

	result := bank.FindNoOfRedistributionCycles(bankValues)
	fmt.Printf("Result : `%d`\n", result)

	result = bank.FindExtendedNoOfRedistributionCycles(bankValues)
	fmt.Printf("Result : `%d`\n", result)
}
