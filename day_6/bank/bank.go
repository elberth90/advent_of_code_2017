package bank

import (
	"strconv"

	"github.com/segmentio/fasthash/fnv1a"
)

// FindNoOfRedistributionCycles return number of redistribution cycles
func FindNoOfRedistributionCycles(bank []int) int {
	hashList := []uint64{hash(bank)}
	for {
		currentElement := getMaxIndex(bank)
		value := bank[currentElement]
		bank[currentElement] = 0

		for i := 0; i < value; i++ {
			currentElement++
			if currentElement > len(bank)-1 {
				currentElement = 0
			}
			bank[currentElement]++
		}

		hash := hash(bank)
		for _, h := range hashList {
			if h == hash {
				return len(hashList)
			}
		}
		hashList = append(hashList, hash)
	}
}

// FindExtendedNoOfRedistributionCycles return number of cycles between second and first occurrence of hash
func FindExtendedNoOfRedistributionCycles(bank []int) int {
	hashList := []uint64{hash(bank)}
	l := 0
	var foundHash uint64
	var found = false
	for {
		currentElement := getMaxIndex(bank)
		value := bank[currentElement]
		bank[currentElement] = 0

		for i := 0; i < value; i++ {
			currentElement++
			if currentElement > len(bank)-1 {
				currentElement = 0
			}
			bank[currentElement]++
		}

		hash := hash(bank)
		if found && hash == foundHash {
			return len(hashList) - l
		}
		if !found {
			for _, h := range hashList {
				if h == hash {
					found = true
					l = len(hashList)
					foundHash = hash
					break
				}
			}
		}
		hashList = append(hashList, hash)
	}
}

func hash(bank []int) uint64 {
	var input string
	for _, b := range bank {
		input += strconv.Itoa(b)
	}
	return fnv1a.HashString64(input)
}

func getMaxIndex(bank []int) int {
	var maxIndex = 0
	var max = bank[maxIndex]

	for index, b := range bank {
		if b > max {
			maxIndex = index
			max = b
		}
	}

	return maxIndex
}
