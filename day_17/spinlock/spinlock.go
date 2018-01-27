package spinlock

const limit = 2017
const largeLimit = 50000000

// PredictLast return number that appears after 2017 value
func PredictLast(steps int) int {
	var buffer = make([]int, limit+1)
	buffer[0] = 0
	var currentPos int

	for i := 1; i <= limit; i++ {
		currentPos = (currentPos+steps)%i + 1
		tmp := []int{}
		tmp = append(tmp, buffer[:currentPos]...)
		tmp = append(tmp, i)
		tmp = append(tmp, buffer[currentPos:len(buffer)-1]...)
		buffer = tmp
	}

	after := 0
	for i := 0; i < len(buffer); i++ {
		if buffer[i] == limit {
			after = buffer[i+1]
			break
		}
	}

	return after
}

// PredictFirst return number that appears after 0 value
func PredictFirst(steps int) int {
	var searchValue int
	var searchPos int
	var afterValue int
	var currentPos int

	for i := 1; i <= largeLimit; i++ {
		currentPos = (currentPos+steps)%i + 1
		if i == searchValue {
			searchPos = currentPos
			continue
		}
		if currentPos <= searchPos {
			searchPos++
		}

		if currentPos == searchPos+1 {
			afterValue = i
		}
	}

	return afterValue
}
