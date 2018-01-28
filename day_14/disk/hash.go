package disk

import (
	"fmt"
	"strconv"
)

const listLength = 256

// GenerateKnotHash generates a Knot Hash
func GenerateKnotHash(input []byte) string {
	inputByte := append(input, []byte{17, 31, 73, 47, 23}...)
	inputLengths := convertToASCIICodes(inputByte)
	sparseHash := generateHash(generateList(), inputLengths, 64)
	denseHash := getDenseHash(sparseHash)
	return convertToHex(denseHash)
}

func generateHash(list [listLength]int, inputLengths []int, repeat int) [listLength]int {
	index, skipSize := 0, 0
	start, end := 0, 0
	for i := 0; i < repeat; i++ {
		for _, length := range inputLengths {
			start = index
			end = (index + length - 1) % len(list)
			list = reverse(list, start, end, length)
			index = (index + length + skipSize) % len(list)
			skipSize++
		}
	}
	return list
}

func reverse(list [listLength]int, start int, end int, length int) [listLength]int {
	for i := 0; i < length/2; i++ {
		start = start % len(list)
		end = end % len(list)
		if end < 0 {
			end = (end + len(list)) % len(list)
		}
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}

	return list
}

func generateList() [listLength]int {
	var list [listLength]int
	for i := 0; i < listLength; i++ {
		list[i] = i
	}

	return list
}

func convertToASCIICodes(list []byte) []int {
	asciiList := make([]int, len(list))
	for i, element := range list {
		asciiList[i] = int(element)
	}

	return asciiList
}

func getDenseHash(sparseHash [listLength]int) []int {
	blockSize := 16
	denseHash := make([]int, blockSize)
	for i := 0; i < blockSize; i++ {
		for _, value := range sparseHash[i*blockSize : (i+1)*blockSize] {
			denseHash[i] ^= value
		}
	}

	return denseHash
}

func convertToHex(list []int) string {
	hexString := ""
	for _, value := range list {
		hexString = fmt.Sprintf("%s%02s", hexString, strconv.FormatInt(int64(value), 16))
	}
	return hexString
}
