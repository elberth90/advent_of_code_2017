package bank

import (
	"testing"
)

func TestFindNoOfRedistributionCycles(t *testing.T) {

	b := NewPredefined(0, 2, 7, 0)
	useCases := map[string]struct {
		bank           *List
		expectedResult int
	}{
		"5 as a result": {
			bank:           b,
			expectedResult: 5,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindNoOfRedistributionCycles(b)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func TestFindExtendedNoOfRedistributionCycles(t *testing.T) {

	b := NewPredefined(0, 2, 7, 0)
	useCases := map[string]struct {
		bank           *List
		expectedResult int
	}{
		"4 as a result": {
			bank:           b,
			expectedResult: 4,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindExtendedNoOfRedistributionCycles(b)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}
