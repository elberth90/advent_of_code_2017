package parser

import (
	"strconv"
	"strings"
	"sync"
	"time"
)

//FindFrequency return frequency of last snd command
func FindFrequency(input string) int {
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")

	registries := map[string]int{}
	var lastSound int

	for i := 0; i < len(instructions); i++ {
		parsedInstruction := strings.Split(instructions[i], " ")
		command := parsedInstruction[0]

		switch command {
		case "snd":
			lastSound = registries[parsedInstruction[1]]
		case "set":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = value
		case "add":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] += value
		case "mul":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = registries[parsedInstruction[1]] * value
		case "mod":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = registries[parsedInstruction[1]] % value
		case "rcv":
			if registries[parsedInstruction[1]] != 0 {
				return lastSound
			}
		case "jgz":
			reg := getValueOrRegistry(parsedInstruction[1], registries[parsedInstruction[1]])
			if reg <= 0 {
				continue
			}

			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			i += value - 1
		}

	}
	return lastSound
}

func getValueOrRegistry(intMaybe string, fallback int) int {
	var value int
	value, err := strconv.Atoi(intMaybe)
	if err != nil {
		value = fallback
	}

	return value
}

// Part2 return how many times process 1 send value to process 0
func Part2(input string) int {
	var wg sync.WaitGroup
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")

	queue1 := make(chan int, 10000)
	queue2 := make(chan int, 10000)
	defer close(queue1)
	defer close(queue2)

	wg.Add(2)

	p0Send := 0
	p1Send := 0
	go program(0, instructions, &p0Send, queue1, queue2, &wg)
	go program(1, instructions, &p1Send, queue2, queue1, &wg)

	wg.Wait()
	return p1Send
}

func program(id int, instructions []string, sendCounter *int, in chan int, out chan int, wg *sync.WaitGroup) {
	registries := map[string]int{"p": id}

	snd := func(value int) bool {
		select {
		case out <- value:
			*sendCounter++
			return false
		case <-time.After(1 * time.Second):
			return true
		}
	}

	rcv := func() (int, bool) {
		select {
		case value := <-in:

			return value, false
		case <-time.After(1 * time.Second):
			return 0, true
		}
	}

	for i := 0; i < len(instructions); i++ {
		parsedInstruction := strings.Split(instructions[i], " ")
		command := parsedInstruction[0]

		switch command {
		case "snd":
			value := getValueOrRegistry(parsedInstruction[1], registries[parsedInstruction[1]])
			timeout := snd(value)
			if timeout {
				wg.Done()
				return
			}

		case "rcv":
			value, timeout := rcv()
			if timeout {
				wg.Done()
				return
			}
			registries[parsedInstruction[1]] = value

		case "set":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = value
		case "add":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] += value
		case "mul":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = registries[parsedInstruction[1]] * value
		case "mod":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = registries[parsedInstruction[1]] % value
		case "jgz":
			reg := getValueOrRegistry(parsedInstruction[1], registries[parsedInstruction[1]])
			if reg <= 0 {
				continue
			}

			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			i += value - 1
		}
	}

	wg.Done()
}
