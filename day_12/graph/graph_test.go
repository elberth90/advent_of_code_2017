package graph

import "testing"

func TestHowManyConnected(t *testing.T) {
	useCases := map[string]struct {
		input         string
		connectedWith string
		result        int
	}{
		"": {
			input:         "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5\n",
			connectedWith: "0",
			result:        6,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := PartOne(uc.input)
			if result != uc.result {
				t.Fatalf("Expected result was `%d`, got `%d`", uc.result, result)
			}
		})
	}
}
