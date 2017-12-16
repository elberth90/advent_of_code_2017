package firewall

import "testing"

func TestFirewall(t *testing.T) {
	useCases := map[string]struct {
		input            string
		expectedSeverity int
	}{
		"24 as a severity": {
			input:            "0: 3\n1: 2\n4: 4\n6: 4",
			expectedSeverity: 24,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := CaughtHowManyTimes(uc.input)
			if result != uc.expectedSeverity {
				t.Fatalf("Expected severity was `%d`, got `%d`", uc.expectedSeverity, result)
			}

		})
	}
}

func TestDonNotGetCaught(t *testing.T) {
	useCases := map[string]struct {
		input         string
		expectedDelay int
	}{
		"10 as a severity": {
			input:         "0: 3\n1: 2\n4: 4\n6: 4",
			expectedDelay: 10,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			result := DonNotGetCaught(uc.input)
			if result != uc.expectedDelay {
				t.Fatalf("Expected delay was `%d`, got `%d`", uc.expectedDelay, result)
			}

		})
	}
}
