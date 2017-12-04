package checksum

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var errParse = errors.New("parse error")

// CalculateChecksum return checksum based on difference between highest and lowest value in row
func CalculateChecksum(input string) (int, error) {
	var sum int = 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		highest, lowest := MinInt, MaxInt
		numbers := strings.Fields(scanner.Text())
		for _, number := range numbers {
			n, err := parseInt(number)
			if err != nil {
				return 0, errors.Wrap(errParse, err.Error())
			}

			if n > highest {
				highest = n
			}

			if n < lowest {
				lowest = n
			}
		}
		sum += highest - lowest
	}
	return sum, nil
}

// CalculateChecksumWithModulo return checksum based on modulo operation between two numbers in row
func CalculateChecksumWithModulo(input string) (int, error) {
	var sum int = 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		highest, lowest := MinInt, MaxInt
		numbers, err := convert(strings.Fields(scanner.Text()))
		if err != nil {
			return 0, err
		}

		for i := 0; i < len(numbers); i++ {
			for j := 0; j < len(numbers); j++ {
				if i == j {
					continue
				}
				n := numbers[i]
				d := numbers[j]
				if n%d == 0 {
					if n > d {
						highest = n
						lowest = d
					} else {
						highest = d
						lowest = n
					}
					break
				}
			}
		}

		sum += highest / lowest
	}
	return sum, nil
}

func parseInt(char string) (int, error) {
	return strconv.Atoi(char)
}

func convert(numbers []string) ([]int, error) {
	var result []int
	for _, number := range numbers {
		n, err := parseInt(number)
		if err != nil {
			return nil, errors.Wrap(errParse, err.Error())
		}
		result = append(result, n)
	}

	return result, nil
}
