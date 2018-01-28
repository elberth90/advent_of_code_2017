package bank

import (
	"testing"
)

func TestFindNoOfRedistributionCycles(t *testing.T) {
	useCases := map[string]struct {
		bank           []int
		expectedResult int
	}{
		"5 as a result": {
			bank:           []int{0, 2, 7, 0},
			expectedResult: 5,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindNoOfRedistributionCycles(uc.bank)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func BenchmarkFindNoOfRedistributionCycles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindNoOfRedistributionCycles([]int{0, 2, 7, 0})
	}
}

func TestFindExtendedNoOfRedistributionCycles(t *testing.T) {

	useCases := map[string]struct {
		bank           []int
		expectedResult int
	}{
		"4 as a result": {
			bank:           []int{0, 2, 7, 0},
			expectedResult: 4,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := FindExtendedNoOfRedistributionCycles(uc.bank)
			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}
