package bridge

import "testing"

func TestLongestBridge(t *testing.T) {
	useCases := map[string]struct {
		elements string
		longest  int
	}{
		"31 as a longest bridge": {
			elements: "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10",
			longest:  31,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			longest := StrongestBridge(uc.elements)
			if longest != uc.longest {
				t.Fatalf("Expected longest bridge was `%d`, got `%d`", uc.longest, longest)
			}
		})
	}
}
