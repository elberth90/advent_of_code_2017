package main

import (
	"bufio"
	"os"

	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_4/validator"
)

const filename = "data.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	counterSimple, counterComplex := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		if validator.SimpleValidate(t) {
			counterSimple += 1
		}
		if validator.ComplexValidate(t) {
			counterComplex += 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", counterSimple)
	fmt.Printf("Result: %d\n", counterComplex)
}
