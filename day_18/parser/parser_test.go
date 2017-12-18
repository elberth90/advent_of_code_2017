package parser

import "testing"

func TestFindFrequency(t *testing.T) {
	useCases := map[string]struct {
		instructions      string
		expectedFrequency int
	}{
		"4 as a result": {
			instructions:      "set a 1\nadd a 2\nmul a a\nmod a 5\nsnd a\nset a 0\nrcv a\njgz a -1\nset a 1\njgz a -2",
			expectedFrequency: 4,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			frequency := FindFrequency(uc.instructions)
			if frequency != uc.expectedFrequency {
				t.Fatalf("Expected frequency was `%d`, got `%d`", uc.expectedFrequency, frequency)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	useCases := map[string]struct {
		instructions string
		snedTimes    int
	}{
		"4 as a result": {
			instructions: "snd 1\nsnd 2\nsnd p\nrcv a\nrcv b\nrcv c\nrcv d",
			snedTimes:    3,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			snedTimes := Part2(uc.instructions)
			if snedTimes != uc.snedTimes {
				t.Fatalf("Expected snedTimes was `%d`, got `%d`", uc.snedTimes, snedTimes)
			}
		})
	}
}
