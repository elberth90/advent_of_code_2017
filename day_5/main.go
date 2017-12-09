package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"

	"github.com/elberth90/advent_of_code_2017/day_5/maze"
)

const filename = "data.txt"

func main() {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	l1 := maze.New()
	l2 := maze.New()
	for scanner.Scan() {
		t := scanner.Text()
		val, err := strconv.Atoi(t)
		if err != nil {
			panic(err)
		}
		l1.Push(val)
		l2.Push(val)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	result := maze.FindExit(l1)
	fmt.Printf("Result: %d\n", result)
	result = maze.FindStrangeExit(l2)
	fmt.Printf("Result: %d\n", result)
}
