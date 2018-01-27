package state_machine

import (
	"github.com/elberth90/advent_of_code_2017/day_25/parser"
)

func Diagnostic(blueprint parser.Blueprint) int {
	var tape []int = make([]int, blueprint.Steps*2, blueprint.Steps*2)

	pointer := blueprint.Steps
	var ins parser.Instruction
	for i, currentState := 0, blueprint.StateList[blueprint.StartWithState]; i < blueprint.Steps; i++ {
		value := tape[pointer]
		//fmt.Printf("%d %#v %#v \n", pointer, currentState, tape)
		if value == 0 {
			ins = currentState.On0
		} else {
			ins = currentState.On1
		}
		//fmt.Printf("%#v \n", ins)
		tape[pointer] = ins.ToWrite
		pointer += ins.Move
		if pointer == len(tape) {
			tape = append(tape, 0)
		} else if pointer == -1 {
			tape = append([]int{0}, tape...)
			pointer = 0
		}
		currentState = blueprint.StateList[ins.NextStateID]
	}

	sum := 0
	for _, v := range tape {
		sum += v
	}

	return sum
}
