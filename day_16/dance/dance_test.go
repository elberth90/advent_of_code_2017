package dance

import "testing"

func TestDance(t *testing.T) {
	useCases := map[string]struct {
		instructions string
		howMany      int
		result       string
	}{
		"aoc test case": {
			instructions: "s1, x3/4, pe/b",
			howMany:      5,
			result:       "baedc",
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := OneTime(uc.howMany, uc.instructions)
			if result != uc.result {
				t.Fatalf("Expected result was `%s`, got `%s`", uc.result, result)
			}
		})
	}
}

func TestDanceNTimes(t *testing.T) {
	useCases := map[string]struct {
		instructions string
		howMany      int
		times        int
		result       string
	}{
		"aoc test case 2 times": {
			instructions: "s1, x3/4, pe/b",
			howMany:      5,
			times:        2,
			result:       "abcde",
		},
		"aoc test case 1000 times": {
			instructions: "s1, x3/4, pe/b",
			howMany:      5,
			times:        9605,
			result:       "baedc",
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := NTimes(uc.howMany, uc.instructions, uc.times)
			if result != uc.result {
				t.Fatalf("Expected result was `%s`, got `%s`", uc.result, result)
			}
		})
	}
}
