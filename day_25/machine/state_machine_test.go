package machine

import (
	"testing"

	"github.com/elberth90/advent_of_code_2017/day_25/parser"
)

func TestDiagnostic(t *testing.T) {
	useCases := map[string]struct {
		blueprint parser.Blueprint
		checksum  int
	}{
		"": {
			blueprint: parser.Blueprint{
				StartWithState: "A",
				Steps:          6,
				StateList: parser.StateList{
					"A": {
						ID: "A",
						On0: parser.Instruction{
							ToWrite:     1,
							Move:        1,
							NextStateID: "B",
						},
						On1: parser.Instruction{
							ToWrite:     0,
							Move:        -1,
							NextStateID: "B",
						},
					},
					"B": {
						ID: "B",
						On0: parser.Instruction{
							ToWrite:     1,
							Move:        -1,
							NextStateID: "A",
						},
						On1: parser.Instruction{
							ToWrite:     1,
							Move:        1,
							NextStateID: "A",
						},
					},
				},
			},
			checksum: 3,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			checksum := Diagnostic(uc.blueprint)
			if checksum != uc.checksum {
				t.Fatalf("Expected checksum was `%d`, got `%d`", uc.checksum, checksum)
			}
		})
	}
}
