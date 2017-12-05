package maze

import (
	"testing"
)

func TestFindExit(t *testing.T) {

	useCases := map[string]struct {
		input          *List
		expectedResult int
	}{
		"exit in 5 steps": {
			input:          NewPredefined(0, 3, 0, 1, -3),
			expectedResult: 5,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindExit(uc.input)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func TestFindStrangeExit(t *testing.T) {

	useCases := map[string]struct {
		input          *List
		expectedResult int
	}{
		"exit in 10 steps": {
			input:          NewPredefined(0, 3, 0, 1, -3),
			expectedResult: 10,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindStrangeExit(uc.input)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}
