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
	l1 := bank.New()
	l2 := bank.New()
	for _, s := range inputList {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		l1.PushFront(i)
		l2.PushFront(i)
	}

	result := bank.FindNoOfRedistributionCycles(l1)
	fmt.Printf("Result : `%d`\n", result)

	result = bank.FindExtendedNoOfRedistributionCycles(l2)
	fmt.Printf("Result : `%d`\n", result)
}
