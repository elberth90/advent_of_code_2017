package main

import (
	"bufio"
	"fmt"
	"os"

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
			counterSimple++
		}
		if validator.ComplexValidate(t) {
			counterComplex++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d\n", counterSimple)
	fmt.Printf("Result: %d\n", counterComplex)
}
