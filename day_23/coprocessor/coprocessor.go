package coprocessor

import (
	"strconv"
	"strings"
)

//HowManyMul return how many mul operations were made
func HowManyMul(input string) int {
	instructions := strings.Split(strings.Trim(input, "\n"), "\n")

	registries := setupRegisters()
	var mulCounter int = 0

	for i := 0; i < len(instructions); i++ {
		parsedInstruction := strings.Split(instructions[i], " ")
		command := parsedInstruction[0]

		switch command {
		case "set":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = value
		case "sub":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] -= value
		case "mul":
			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			registries[parsedInstruction[1]] = registries[parsedInstruction[1]] * value
			mulCounter++
		case "jnz":
			reg := getValueOrRegistry(parsedInstruction[1], registries[parsedInstruction[1]])
			if reg == 0 {
				continue
			}

			value := getValueOrRegistry(parsedInstruction[2], registries[parsedInstruction[2]])
			i += value - 1
		}

	}
	return mulCounter
}

// HRegisterValue return value from H register
func HRegisterValue() int {
	registers := setupRegisters()
	registers["a"] = 1
	registers["b"] = 79
	registers["c"] = 79

	registers["b"] = registers["b"] * 100
	registers["b"] += 100000
	registers["c"] = registers["b"]
	registers["c"] += 17000

	for ok := true; ok; ok = registers["g"] != 0 {
		registers["f"] = 1
		registers["d"] = 2

		for d := registers["d"]; d*d < registers["b"]; d++ {
			if registers["b"]%d == 0 {
				registers["f"] = 0
				break
			}
		}
		if registers["f"] == 0 {
			registers["h"]++
		}
		registers["g"] = registers["b"]
		registers["g"] -= registers["c"]
		registers["b"] += 17
	}
	return registers["h"]
}

func getValueOrRegistry(intMaybe string, fallback int) int {
	var value int
	value, err := strconv.Atoi(intMaybe)
	if err != nil {
		value = fallback
	}

	return value
}

func setupRegisters() map[string]int {
	var registers map[string]int = map[string]int{}
	for i := 97; i < 105; i++ {
		registers[string(i)] = 0
	}

	return registers
}
