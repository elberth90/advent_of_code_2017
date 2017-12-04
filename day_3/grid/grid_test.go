package grid

import "testing"

func TestCalculateSteps(t *testing.T) {
	useCases := map[string]struct {
		gridNo              int
		expectedStepsResult int
	}{
		"0 as a result": {
			gridNo:              1,
			expectedStepsResult: 0,
		},
		"3 as a result": {
			gridNo:              12,
			expectedStepsResult: 3,
		},
		"2 as a result": {
			gridNo:              23,
			expectedStepsResult: 2,
		},
		"31 as a result": {
			gridNo:              1024,
			expectedStepsResult: 31,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			steps := CalculateSteps(uc.gridNo)
			if steps != uc.expectedStepsResult {
				t.Fatalf("Expetced number of steps was `%d`, got `%d`", uc.expectedStepsResult, steps)
			}
		})
	}
}
