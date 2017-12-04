package captcha

import (
	"strconv"

	"github.com/pkg/errors"
)

var errParse = errors.New("parse error")

// CalculateSum calculate captcha sum  with jump equal to 1
func CalculateSum(input string) (int, error) {
	return calculate(input, 1)
}

// CalculateSumWithJump calculate captcha sum with jump equal to half of the input
func CalculateSumWithJump(input string) (int, error) {
	jump := len(input) / 2
	return calculate(input, jump)
}

func calculate(input string, jump int) (int, error) {
	var sum int = 0
	var next int = 0
	inputSize := len(input)

	for i := 0; i < inputSize; i += 1 {
		current, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return 0, errors.Wrap(errParse, err.Error())
		}

		nextElementIndex := (i + jump) % inputSize
		next, err = strconv.Atoi(string(input[nextElementIndex]))
		if err != nil {
			return 0, errors.Wrap(errParse, err.Error())
		}

		if current == next {
			sum += current
		}
	}

	return sum, nil
}
