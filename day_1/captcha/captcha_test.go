package captcha

import (
	"testing"

	"github.com/pkg/errors"
)

func TestCalculateSum(t *testing.T) {
	useCases := map[string]struct {
		input          string
		expectedResult int
		err            error
	}{
		"not int": {
			input: "1122e33",
			err:   errParse,
		},
		"3 as a result": {
			input:          "1122",
			expectedResult: 3,
		},
		"4 as a result": {
			input:          "1111",
			expectedResult: 4,
		},
		"0 as a result": {
			input:          "1234",
			expectedResult: 0,
		},
		"9 as a result": {
			input:          "91212129",
			expectedResult: 9,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := CalculateSum(uc.input)
			if errors.Cause(err) != uc.err {
				t.Fatalf("Expected error was `%s`, got `%s`", uc.err, err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func TestCalculateSumWithOverhead(t *testing.T) {
	useCases := map[string]struct {
		input          string
		expectedResult int
		err            error
	}{
		"not int": {
			input: "1122e33",
			err:   errParse,
		},
		"6 as a result of 1212": {
			input:          "1212",
			expectedResult: 6,
		},
		"0 as a result of 1221": {
			input:          "1221",
			expectedResult: 0,
		},
		"4 as a result of 123425": {
			input:          "123425",
			expectedResult: 4,
		},
		"12 as a result of 123123": {
			input:          "123123",
			expectedResult: 12,
		},
		"4 as a result of 12131415": {
			input:          "12131415",
			expectedResult: 4,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := CalculateSumWithJump(uc.input)
			if errors.Cause(err) != uc.err {
				t.Fatalf("Expected error was `%s`, got `%s`", uc.err, err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}
