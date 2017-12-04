package checksum

import (
	"testing"

	"github.com/pkg/errors"
)

func TestCalculateChecksum(t *testing.T) {
	useCases := map[string]struct {
		input          string
		expectedResult int
		err            error
	}{
		"not int": {
			input: "5 1 9 5\n7 5 g\na 4 6 8\n",
			err:   errParse,
		},
		"expected checksum 18": {
			input:          "5 1 9 5\n7 5 3\n2 4 6 8\n",
			expectedResult: 18,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := CalculateChecksum(uc.input)
			if errors.Cause(err) != uc.err {
				t.Fatalf("Expected error was `%s`, got `%s`", uc.err, err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func TestCalculateChecksumWithModulo(t *testing.T) {
	useCases := map[string]struct {
		input          string
		expectedResult int
		err            error
	}{
		//"not int": {
		//	input: "5 1 9 5\n7 5 g\na 4 6 8\n",
		//	err:   errParse,
		//},
		"expected checksum 18": {
			input:          "5 9 2 8\n9 4 7 3\n3 8 6 5\n",
			expectedResult: 9,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := CalculateChecksumWithModulo(uc.input)
			if errors.Cause(err) != uc.err {
				t.Fatalf("Expected error was `%s`, got `%s`", uc.err, err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}
