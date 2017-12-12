package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	useCases := map[string]struct {
		listLength int
		length     []int
		result     int
	}{
		"12 as a result": {
			listLength: 5,
			length:     []int{3, 4, 1, 5},
			result:     12,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := Part1(uc.listLength, uc.length)
			if result != uc.result {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.result, result)
			}

		})
	}
}
