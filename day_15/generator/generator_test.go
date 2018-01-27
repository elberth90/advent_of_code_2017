package generator

import "testing"

func TestJudgeCount(t *testing.T) {
	useCases := map[string]struct {
		firstInput    int64
		secondInput   int64
		expectedCount int
	}{
		"588 as result": {
			firstInput:    65,
			secondInput:   8921,
			expectedCount: 588,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := JudgeCount(uc.firstInput, uc.secondInput)
			if result != uc.expectedCount {
				t.Fatalf("Expected count was `%d`, got `%d`", uc.expectedCount, result)
			}
		})
	}
}

func TestExtendedJudgeCount(t *testing.T) {
	useCases := map[string]struct {
		firstInput    int64
		secondInput   int64
		expectedCount int
	}{
		"588 as result": {
			firstInput:    65,
			secondInput:   8921,
			expectedCount: 309,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := ExtendedJudgeCount(uc.firstInput, uc.secondInput)
			if result != uc.expectedCount {
				t.Fatalf("Expected count was `%d`, got `%d`", uc.expectedCount, result)
			}
		})
	}
}
