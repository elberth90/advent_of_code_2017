package firewall

import (
	"strconv"
	"strings"
)

// CaughtHowManyTimes return amount of time program was caught by firewall
func CaughtHowManyTimes(input string) int {
	data := strings.Split(input, "\n")
	severity := calculateSeverity(data)
	return severity
}

func calculateSeverity(data []string) int {

	var fList = map[int]int{}

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
		fList[index] = weight
	}

	severity := 0
	for index, weight := range fList {
		if weight == 1 || index%(2*(weight-1)) == 0 {
			severity += index * weight
		}
	}

	return severity
}

func calculateDelay(data []string) int {

	var fList = map[int]int{}

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
		fList[index] = weight
	}

	delay := 0
	isCaught := true
	for isCaught {
		isCaught = false
		for index, weight := range fList {
			if weight == 1 || (index+delay)%(2*(weight-1)) == 0 {
				isCaught = true
				delay++
				break
			}
		}
	}
	return delay
}

// DonNotGetCaught return delay that should be applied to not be caught by firewall
func DonNotGetCaught(input string) int {
	data := strings.Split(input, "\n")
	return calculateDelay(data)
}
