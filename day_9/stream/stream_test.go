package stream

import "testing"

func TestStream(t *testing.T) {
	useCases := map[string]struct {
		inputStream  string
		totalScore   int
		garbageCount int
	}{
		"score of 1": {
			inputStream:  "{}",
			totalScore:   1,
			garbageCount: 0,
		},
		"score of 6": {
			inputStream:  "{{{}}}",
			totalScore:   6,
			garbageCount: 0,
		},
		"score of 5": {
			inputStream:  "{{},{}}",
			totalScore:   5,
			garbageCount: 0,
		},
		"score of 16": {
			inputStream:  "{{{},{},{{}}}}",
			totalScore:   16,
			garbageCount: 0,
		},
		"score of 1 again": {
			inputStream:  "{<a>,<a>,<a>,<a>}",
			totalScore:   1,
			garbageCount: 4,
		},
		"score of 9": {
			inputStream:  "{{<ab>},{<ab>},{<ab>},{<ab>}}",
			totalScore:   9,
			garbageCount: 8,
		},
		"score of 9 again": {
			inputStream:  "{{<!!>},{<!!>},{<!!>},{<!!>}}",
			totalScore:   9,
			garbageCount: 0,
		},
		"score of 3": {
			inputStream:  "{{<a!>},{<a!>},{<a!>},{<ab>}}",
			totalScore:   3,
			garbageCount: 17,
		},
		"bla": {
			inputStream:  `!> !}`,
			totalScore:   0,
			garbageCount: 0,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			score, garbageCount := CountScore(uc.inputStream)
			if score != uc.totalScore {
				t.Fatalf("Expected score was `%d`, got `%d`", uc.totalScore, score)
			}
			if garbageCount != uc.garbageCount {
				t.Fatalf("Expected garbage count was `%d`, got `%d`", uc.garbageCount, garbageCount)
			}
		})
	}
}
