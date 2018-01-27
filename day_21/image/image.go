package image

import (
	"fmt"
	"strings"
)

// MakeArt apply rules into to the image in n iterations
func MakeArt(rules string, iterations int) int {
	r := parseRules(rules)

	image := []string{
		".#.",
		"..#",
		"###",
	}

	var step int
	for i := 0; i < iterations; i++ {
		newImage := []string{}
		size := len(image)
		if size%2 == 0 {
			step = 2
		} else {
			step = 3
		}

		for j := 0; j < size; j += step {
			for k := 0; k < step+1; k++ {
				newImage = append(newImage, "")
			}

			for l := 0; l < size; l += step {
				block := []string{}
				for m := 0; m < step; m++ {
					block = append(block, image[j+m][l:l+step])
				}
				block = transform(block, r)
				dstY := (j / step) * (step + 1)
				for m := 0; m < step+1; m++ {
					newImage[dstY+m] += block[m]
				}
			}
		}
		image = newImage
	}

	var counter = 0
	for _, r := range image {
		counter += strings.Count(r, "#")
	}
	return counter
}

func transform(block []string, rules map[string]string) []string {
	applyRule := func(block []string, rules map[string]string) ([]string, bool) {
		b := strings.Join(block, "/")
		for index, r := range rules {
			if index == b {
				return strings.Split(r, "/"), true
			}
		}
		return block, false
	}

	for i := 0; i < 4; i++ {
		block = rotate(block)
		if block, ok := applyRule(block, rules); ok {
			return block
		}

		flippedBlock := flip(block)
		if block, ok := applyRule(flippedBlock, rules); ok {
			return block
		}
	}

	return block
}

func rotate(block []string) []string {
	newBlock := make([]string, len(block))
	n := len(block)
	for i := 0; i < n; i++ {
		var e string
		for j := 0; j < n; j++ {
			e = fmt.Sprintf("%s%s", e, string(block[n-j-1][i]))
		}
		newBlock[i] = e
	}

	return newBlock
}

func flip(block []string) []string {
	reverse := func(s string) string {
		var result string
		for _, v := range s {
			result = string(v) + result
		}
		return result
	}

	var newBlock = make([]string, len(block))
	for i, b := range block {
		newBlock[i] = reverse(b)
	}

	return newBlock
}

func parseRules(rules string) map[string]string {
	ret := map[string]string{}
	ruleList := strings.Split(strings.Trim(rules, "\n"), "\n")
	for _, r := range ruleList {
		d := strings.Split(r, " => ")
		ret[d[0]] = d[1]
	}

	return ret
}
