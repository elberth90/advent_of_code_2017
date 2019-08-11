package dance

import (
	"reflect"
	"strconv"
	"strings"
)

const a = 97

// OneTime perform dance only one times
func OneTime(howMany int, instructions string) string {
	line := make([]string, howMany)
	for i, j := a, 0; i < a+howMany; i++ {
		line[j] = string(i)
		j++
	}

	line = move(line, instructions)

	return strings.Join(line, "")
}

func move(line []string, instructions string) []string {
	for _, i := range strings.Split(instructions, ",") {
		i = strings.TrimSpace(i)
		switch string([]rune(i)[0]) {
		case "s":
			j, err := strconv.Atoi(string([]rune(i)[1:]))
			if err != nil {
				panic(err)
			}
			line = spin(line, j)
		case "x":
			j := strings.Split(string([]rune(i)[1:]), "/")
			first, err := strconv.Atoi(j[0])
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(j[1])
			if err != nil {
				panic(err)
			}
			line = exchange(line, first, second)
		case "p":
			j := strings.Split(string([]rune(i)[1:]), "/")
			line = partner(line, j[0], j[1])
		default:
			panic(string([]rune(i)[0]))
		}
	}

	return line
}

func spin(list []string, size int) []string {
	l := make([]string, 0, len(list))
	l = append(l, list[len(list)-size:]...)
	l = append(l, list[:len(list)-size]...)
	return l
}

func exchange(list []string, first int, second int) []string {
	tmp := list[second]
	list[second] = list[first]
	list[first] = tmp
	return list
}

func partner(list []string, first string, second string) []string {
	var firstPosition int
	var secondPosition int
	for i, value := range list {
		if value == first {
			firstPosition = i
		}
		if value == second {
			secondPosition = i
		}
	}

	return exchange(list, firstPosition, secondPosition)
}

// NTimes perform dance N times
func NTimes(howMany int, instructions string, times int) string {
	line := make([]string, howMany)
	initial := line
	for i, j := a, 0; i < a+howMany; i++ {
		line[j] = string(i)
		j++
	}

	var cycle int
	for i := 0; i < times; i++ {
		line = move(line, instructions)
		if reflect.DeepEqual(initial, line) {
			cycle = i + 1
			break
		}
	}

	line = initial
	if cycle != 0 {
		for i := 0; i < (times % cycle); i++ {
			line = move(line, instructions)
		}
	}

	return strings.Join(line, "")
}
