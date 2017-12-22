package virus

import "testing"

func TestVirusInfections(t *testing.T) {
	useCases := map[string]struct {
		initialGrid string
		bursts      int
		infections  int
	}{
		"5 infections during 7 bursts": {
			initialGrid: "..#\n#..\n...\n",
			bursts:      7,
			infections:  5,
		},
		"41 infections during 70 bursts": {
			initialGrid: "..#\n#..\n...\n",
			bursts:      70,
			infections:  41,
		},
		"5587 infections during 10000 bursts": {
			initialGrid: "..#\n#..\n...\n",
			bursts:      10000,
			infections:  5587,
		},
	}
	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			infections := CalculateInfections(uc.initialGrid, uc.bursts)
			if infections != uc.infections {
				t.Fatalf("Expected number of infections was `%d`, got `%d`", infections, uc.infections)
			}
		})
	}
}

func TestCalculateInfectionsWithSpeedUp(t *testing.T) {
	useCases := map[string]struct {
		initialGrid string
		bursts      int
		infections  int
	}{
		"26 infections during 100 bursts": {
			initialGrid: "..#\n#..\n...\n",
			bursts:      100,
			infections:  26,
		},
		"2511944 infections during 10000000 bursts": {
			initialGrid: "..#\n#..\n...\n",
			bursts:      10000000,
			infections:  2511944,
		},
	}
	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			infections := CalculateInfectionsWithSpeedUp(uc.initialGrid, uc.bursts)
			if infections != uc.infections {
				t.Fatalf("Expected number of infections was `%d`, got `%d`", uc.infections, infections)
			}
		})
	}
}
