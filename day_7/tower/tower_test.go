package tower

import "testing"

var input = "pbga (66)\nxhth (57)\nebii (61)\nhavc (66)\nktlj (57)\nfwft (72) -> ktlj, cntj, xhth\nqoyq (66)\npadx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)\nugml (68) -> gyxo, ebii, jptl\ngyxo (61)\ncntj (57)"

func TestFindBottomProgram(t *testing.T) {
	useCases := map[string]struct {
		input           string
		expectedProgram string
	}{
		"tknk as result": {
			input:           input,
			expectedProgram: "tknk",
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			program := FindBottomProgram(uc.input)
			if program != uc.expectedProgram {
				t.Fatalf("Expected program was `%s`, got `%s`", uc.expectedProgram, program)
			}
		})
	}
}

func TestGetWeightToBalance(t *testing.T) {
	useCases := map[string]struct {
		input          string
		expectedWeight int
	}{
		"9  as result": {
			input:          "a (1) -> b, c, d\nb (7) -> e, f\nc (3) -> g, h, i\nd (2)\ne (1)\nf (1)\ng (2)\nh (2)\ni (2)",
			expectedWeight: 9,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			weight := GetWeightToBalance(uc.input)
			if weight != uc.expectedWeight {
				t.Fatalf("Expected waight was `%d`, got `%d`", uc.expectedWeight, weight)
			}
		})
	}
}
