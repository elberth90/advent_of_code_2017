package bank

// FindNoOfRedistributionCycles return number of redistribution cycles
func FindNoOfRedistributionCycles(banks *List) int {
	hashList := []string{banks.Hash()}
	for {
		currentElement := banks.Max()
		value := currentElement.Value
		currentElement.Value = 0

		for i := 0; i < value; i++ {
			currentElement = currentElement.Prev()
			currentElement.Value++
		}

		hash := banks.Hash()
		for _, h := range hashList {
			if h == hash {
				return len(hashList)
			}
		}
		hashList = append(hashList, hash)
	}
}

// FindExtendedNoOfRedistributionCycles return number of cycles between second and first occurrence of hash
func FindExtendedNoOfRedistributionCycles(banks *List) int {
	hashList := []string{banks.Hash()}
	l := 0
	foundHash := ""
	for {
		currentElement := banks.Max()
		value := currentElement.Value
		currentElement.Value = 0

		for i := 0; i < value; i++ {
			currentElement = currentElement.Prev()
			currentElement.Value++
		}

		hash := banks.Hash()
		if hash == foundHash {
			return len(hashList) - l
		}
		if foundHash == "" {
			for _, h := range hashList {
				if h == hash {
					l = len(hashList)
					foundHash = hash
					break
				}
			}
		}
		hashList = append(hashList, hash)
	}
}
