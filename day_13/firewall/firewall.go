package firewall

import (
	"strconv"
	"strings"
)

type layer struct {
	index  int
	weight int
}

type firewall []layer

// CaughtHowManyTimes return amount of time program was caught by firewall
func CaughtHowManyTimes(input string) int {
	data := strings.Split(input, "\n")
	severity := calculateSeverity(data)
	return severity
}

// DonNotGetCaught return delay that should be applied to not be caught by firewall
func DonNotGetCaught(input string) int {
	data := strings.Split(input, "\n")
	return calculateDelay(data)
}

func calculateSeverity(data []string) int {

	firewall := make(firewall, len(data))
	var i = 0
	for _, line := range data {
		l := strings.Split(line, ": ")
		index, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}
		weight, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}
		firewall[i] = layer{index: index, weight: weight}
		i++
	}

	severity := 0
	for _, layer := range firewall {
		if layer.weight == 1 || layer.index%(2*(layer.weight-1)) == 0 {
			severity += layer.index * layer.weight
		}
	}

	return severity
}

func calculateDelay(data []string) int {

	firewall := make(firewall, len(data))
	var i = 0
	for _, line := range data {
		l := strings.Split(line, ": ")
		index, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}
		weight, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}
		firewall[i] = layer{index: index, weight: weight}
		i++
	}

	var delay = 0
	isCaught := true
	for isCaught {
		isCaught = false
		for _, layer := range firewall {
			if layer.weight == 1 || (layer.index+delay)%(2*(layer.weight-1)) == 0 {
				isCaught = true
				delay++
				break
			}
		}
	}
	return delay
}
