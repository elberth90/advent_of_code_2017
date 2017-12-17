package spinlock

import "testing"

func TestPredict(t *testing.T) {
	useCases := map[string]struct {
		steps int
		after int
	}{
		"with step 3": {
			steps: 3,
			after: 638,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			after := PredictLast(uc.steps)
			if after != uc.after {
				t.Fatalf("Expected after was `%d`, got `%d`", uc.after, after)
			}
		})
	}
}
