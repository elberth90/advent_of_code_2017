package disk

import (
	"fmt"
	"strconv"
)

// Part1 return result of multiplication of the first numbers from generate hash
func Part1(listLength int, inputLengths []int) int {
	list := generateList(listLength)
	h := generateHash(list, inputLengths, 1)

	return h[0] * h[1]
}

// Part2 generates a Knot Hash
func Part2(listLength int, input []byte) string {
	inputByte := append(input, []byte{17, 31, 73, 47, 23}...)
	inputLengths := convertToASCIICodes(inputByte)
	sparseHash := generateHash(generateList(listLength), inputLengths, 64)
	denseHash := getDenseHash(sparseHash)
	return convertToHex(denseHash)
}

func generateHash(list []int, inputLengths []int, repeat int) []int {
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

func reverse(list []int, start int, end int, length int) []int {
	copyList := make([]int, len(list))
	copy(copyList, list)

	for i := 0; i < length; i++ {
		start = start % len(list)
		end = end % len(list)
		if end < 0 {
			end = (end + len(list)) % len(list)
		}
		list[start] = copyList[end]
		start++
		end--
	}

	return list
}

func generateList(size int) []int {
	list := make([]int, size)
	for i := 0; i < size; i++ {
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

func getDenseHash(sparseHash []int) []int {
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
