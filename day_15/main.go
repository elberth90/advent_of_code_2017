package main

import (
	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_15/generator"
)

func main() {

	result := generator.JudgeCount(int64(679), int64(771))
	fmt.Printf("Result: `%d`\n", result)

	result = generator.ExtendedJudgeCount(int64(679), int64(771))
	fmt.Printf("Result: `%d`\n", result)
}
