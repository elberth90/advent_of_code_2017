package parser

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	startRegex = `Begin in state (?P<start>\w+)`
	stepsRegex = `Perform a diagnostic checksum after (?P<Steps>\d+)`
)

// Blueprint contain information about blueprint
type Blueprint struct {
	StartWithState string
	Steps          int
	StateList      StateList
}

// State contain information about state
type State struct {
	ID  string
	On0 Instruction
	On1 Instruction
}

// Instruction contain instructions for state
type Instruction struct {
	ToWrite     int
	Move        int
	NextStateID string
}

// StateList holds map of the State
type StateList = map[string]State

// ParseBlueprint return Blueprint struct
func ParseBlueprint(input string) Blueprint {
	var err error
	b := Blueprint{
		StateList: StateList{},
	}

	r := regexp.MustCompile(startRegex)
	b.StartWithState = r.FindStringSubmatch(input)[1]

	r = regexp.MustCompile(stepsRegex)
	b.Steps, err = strconv.Atoi(r.FindStringSubmatch(input)[1])
	if err != nil {
		panic(err)
	}

	data := strings.Split(strings.Trim(input, "\n"), "\n")

	for _, i := range split(data[3:], 10) {
		state := build(i)
		b.StateList[state.ID] = state
	}

	return b
}

func build(i []string) State {
	var err error

	id := strings.Split(strings.Trim(i[0], ":"), " ")[2]

	on0 := Instruction{}
	on0.ToWrite, err = strconv.Atoi(strings.Split(strings.TrimSpace(strings.Trim(i[2], ".")), " ")[4])
	if err != nil {
		panic(err)
	}
	if strings.Split(strings.TrimSpace(strings.Trim(i[3], ".")), " ")[6] == "right" {
		on0.Move = 1
	} else {
		on0.Move = -1
	}
	on0.NextStateID = strings.Split(strings.TrimSpace(strings.Trim(i[4], ".")), " ")[4]

	on1 := Instruction{}
	on1.ToWrite, err = strconv.Atoi(strings.Split(strings.TrimSpace(strings.Trim(i[6], ".")), " ")[4])
	if err != nil {
		panic(err)
	}
	if strings.Split(strings.TrimSpace(strings.Trim(i[7], ".")), " ")[6] == "right" {
		on1.Move = 1
	} else {
		on1.Move = -1
	}
	on1.NextStateID = strings.Split(strings.TrimSpace(strings.Trim(i[8], ".")), " ")[4]

	return State{ID: id, On0: on0, On1: on1}
}

func split(buf []string, lim int) [][]string {
	var chunk []string
	chunks := make([][]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
