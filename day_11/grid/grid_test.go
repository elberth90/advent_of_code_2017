package grid

import "testing"

func TestGetDistance(t *testing.T) {
	useCases := map[string]struct {
		instructions string
		distance     int
	}{
		"3 as a distance": {
			instructions: "ne,ne,ne\n",
			distance:     3,
		},
		"0 as a distance": {
			instructions: "ne,ne,sw,sw",
			distance:     0,
		},
		"2 as a distance": {
			instructions: "ne,ne,s,s",
			distance:     2,
		},
		"3 again as a distance": {
			instructions: "se,sw,se,sw,sw",
			distance:     3,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			distance := GetDistance(uc.instructions)
			if distance != uc.distance {
				t.Fatalf("Expected distance was `%d`, got `%d`", uc.distance, distance)
			}
		})
	}
}
