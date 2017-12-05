package validator

import (
	"sort"
	"strings"
)

const separator = " "

// SimpleValidate validate given passphrase searching for duplicate words
func SimpleValidate(passphrase string) bool {
	splited := strings.Split(passphrase, separator)

	blockHolder := make(map[string]uint)
	for _, s := range splited {
		if _, ok := blockHolder[s]; ok {
			return false
		}
		blockHolder[s] = 1
	}
	return true
}

// ComplexValidate validate given passphrase searching for duplicate words after rearrange
func ComplexValidate(passphrase string) bool {
	splited := strings.Split(passphrase, separator)

	blockHolder := make(map[string]uint)
	for _, s := range splited {
		sortedS := strings.Split(s, "")
		sort.Strings(sortedS)
		s = strings.Join(sortedS, "")
		if _, ok := blockHolder[s]; ok {
			return false
		}
		blockHolder[s] = 1
	}
	return true
}
