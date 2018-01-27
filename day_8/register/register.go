package register

import (
	"strconv"
	"strings"
)

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

const (
	increase = iota
	decrease
)

// Run is a method that parse instruction and execute them
func Run(input string) (int, int) {
	data := strings.Split(input, "\n")
	m := map[string]int{}
	max := minInt
	highestMax := minInt

	for _, l := range data {
		s := split(l)
		if _, ok := m[s.ElementToChange]; !ok {
			m[s.ElementToChange] = 0
		}
		if _, ok := m[s.StatementElement]; !ok {
			m[s.StatementElement] = 0
		}

		if ifFunc(m[s.StatementElement], s.Statement, s.StatementValue) {
			c := changedValue(m[s.ElementToChange], s.Change, s.ChangeValue)
			m[s.ElementToChange] = c
		}

		if m[s.ElementToChange] > highestMax {
			highestMax = m[s.ElementToChange]
		}
	}

	for _, e := range m {
		if e > max {
			max = e
		}
	}

	return max, highestMax
}

func ifFunc(valueFirst int, statement string, valueSecond int) bool {
	switch statement {
	case "<":
		return valueFirst < valueSecond
	case "<=":
		return valueFirst <= valueSecond
	case ">":
		return valueFirst > valueSecond
	case ">=":
		return valueFirst >= valueSecond
	case "==":
		return valueFirst == valueSecond
	case "!=":
		return valueFirst != valueSecond
	default:
		panic("bad")
	}
}

func changedValue(element int, change int, changeValue int) int {
	switch change {
	case increase:
		return element + changeValue
	case decrease:
		return element - changeValue
	default:
		panic("unsupported operator")
	}
}

type element struct {
	ElementToChange string
	Change          int
	ChangeValue     int

	StatementElement string
	Statement        string
	StatementValue   int
}

func split(line string) *element {
	data := strings.Split(line, " ")

	var change int

	switch data[1] {
	case "inc":
		change = increase
	case "dec":
		change = decrease
	}

	changeValue, err := strconv.Atoi(data[2])
	if err != nil {
		panic(err)
	}

	statementValue, err := strconv.Atoi(data[6])
	if err != nil {
		panic(err)
	}

	return &element{
		ElementToChange: data[0],
		Change:          change,
		ChangeValue:     changeValue,

		StatementElement: data[4],
		Statement:        data[5],
		StatementValue:   statementValue,
	}
}
