package parser

import (
	"fmt"
	"reflect"
	"testing"
)

var input = `Begin in state A.
Perform a diagnostic checksum after 6 Steps.

In State A:
If the current value is 0:
- Write the value 1.
- Move one slot to the right.
- Continue with State B.
If the current value is 1:
- Write the value 0.
- Move one slot to the left.
- Continue with State B.

In State B:
If the current value is 0:
- Write the value 1.
- Move one slot to the left.
- Continue with State A.
If the current value is 1:
- Write the value 1.
- Move one slot to the right.
- Continue with State A.
`

func TestParseInput(t *testing.T) {
	useCases := map[string]struct {
		input     string
		blueprint Blueprint
	}{
		"": {
			input: input,
			blueprint: Blueprint{
				StartWithState: "A",
				Steps:          6,
				StateList: StateList{
					"A": {
						ID: "A",
						On0: Instruction{
							ToWrite:     1,
							Move:        1,
							NextStateID: "B",
						},
						On1: Instruction{
							ToWrite:     0,
							Move:        -1,
							NextStateID: "B",
						},
					},
					"B": {
						ID: "B",
						On0: Instruction{
							ToWrite:     1,
							Move:        -1,
							NextStateID: "A",
						},
						On1: Instruction{
							ToWrite:     1,
							Move:        1,
							NextStateID: "A",
						},
					},
				},
			},
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			blueprint := ParseBlueprint(uc.input)
			if !reflect.DeepEqual(blueprint, uc.blueprint) {
				p(uc.blueprint)
				p(blueprint)
				t.Fail()
			}
		})
	}
}

func p(b Blueprint) {
	fmt.Println()
	fmt.Printf("StartWithState: %s \n", b.StartWithState)
	fmt.Printf("Steps: %d \n", b.Steps)
	for id, s := range b.StateList {
		fmt.Printf("State map ID %s \n", id)
		fmt.Printf("State ID %s \n", s.ID)
		fmt.Printf("State On0 %#+v \n", s.On0)
		fmt.Printf("State On1 %#+v \n", s.On1)
	}
	fmt.Println()
}
