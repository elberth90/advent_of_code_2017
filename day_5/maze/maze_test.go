package maze

import (
	"testing"
)

func TestFindExit(t *testing.T) {

	useCases := map[string]struct {
		input          string
		expectedResult int
	}{
		"exit in 5 steps": {
			input:          "0\n3\n0\n1\n-3\n",
			expectedResult: 5,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := FindExit(uc.input)
			if err != nil {
				t.Fatalf("Unexpected error: `%#v`", err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func TestFindStrangeExit(t *testing.T) {

	useCases := map[string]struct {
		input          string
		expectedResult int
	}{
		"exit in 10 steps": {
			input:          "0\n3\n0\n1\n-3\n",
			expectedResult: 10,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result, err := FindStrangeExit(uc.input)
			if err != nil {
				t.Fatalf("Unexpected error: `%#v`", err)
			}

			if result != uc.expectedResult {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.expectedResult, result)
			}
		})
	}
}

func BenchmarkFindStrangeExit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := FindStrangeExit("0\n3\n0\n1\n-3\n")
		if err != nil {
			b.Error(err)
		}
	}
}
