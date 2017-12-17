package main

import (
	"fmt"

	"github.com/elberth90/advent_of_code_2017/day_17/spinlock"
)

func main() {

	result := spinlock.PredictLast(328)
	fmt.Printf("Result: `%d`\n", result)

	result = spinlock.PredictFirst(328)
	fmt.Printf("Result: `%d`\n", result)
}
