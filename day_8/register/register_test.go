package register

import "testing"

func TestRun(t *testing.T) {
	useCases := map[string]struct {
		input     string
		localMax  int
		globalMax int
	}{
		"ACO test case": {
			input:     "b inc 5 if a > 1\na inc 1 if b < 5\nc dec -10 if a >= 1\nc inc -20 if c == 10",
			localMax:  1,
			globalMax: 10,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			localMax, globalMax := Run(uc.input)
			if localMax != uc.localMax {
				t.Fatalf("Expected local max was `%d`, go `%d`", uc.localMax, localMax)
			}

			if globalMax != uc.globalMax {
				t.Fatalf("Expected gloabl max was `%d`, go `%d`", uc.globalMax, globalMax)
			}
		})
	}
}
