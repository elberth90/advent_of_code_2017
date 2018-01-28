package disk

import "testing"

func TestDiskOccupancy(t *testing.T) {
	useCases := map[string]struct {
		input     string
		occupancy int
	}{
		"AOC test case": {
			input:     "flqrgnkx",
			occupancy: 8108,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			occupancy := Occupancy(uc.input)
			if occupancy != uc.occupancy {
				t.Fatalf("Expected occupancy was `%d`, got `%d`", uc.occupancy, occupancy)
			}
		})
	}
}

func BenchmarkOccupancy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RegionsFinder("flqrgnkx")
	}
}
