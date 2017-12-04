package main

import (
	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_3/grid"
)

const input_1 = 265149

func main() {
	sum := grid.CalculateSteps(input_1)
	fmt.Printf("Result: `%d`\n", sum)

	number := grid.CalculateNext(input_1)
	fmt.Printf("Result: `%d`\n", number)
}
