package generator

const (
	divider         = int64(2147483647)
	genAFactor      = int64(16807)
	genARestriction = int64(4)
	genBFactor      = int64(48271)
	genBRestriction = int64(8)
	basicLimit      = 40000000
	extendedLimit   = 5000000
)

type criteriaFunc func(value int64) bool

// JudgeCount return number of pairs
func JudgeCount(firstGenInput int64, secondGenInput int64) int {
	dummyCriteria := func(value int64) bool {
		return true
	}
	c1 := runGenerator(firstGenInput, genAFactor, dummyCriteria)
	c2 := runGenerator(secondGenInput, genBFactor, dummyCriteria)
	counter := newJudge(basicLimit, c1, c2)
	return counter
}

func newJudge(limit int, recA chan int64, recB chan int64) int {
	var counter int
	for i := 0; i < limit; i++ {
		if <-recA&0xFFFF == <-recB&0xFFFF {
			counter++
		}
	}

	return counter
}

func runGenerator(startingValue int64, factor int64, criteria criteriaFunc) chan int64 {
	c := make(chan int64, 20)
	go func() {
		defer close(c)
		last := startingValue
		for {
			select {
			case c <- last:
				for true {
					last = (last * factor) % divider
					if criteria(last) {
						break
					}
				}
			}
		}
	}()
	return c
}

// ExtendedJudgeCount return number of pairs using different generator logic
func ExtendedJudgeCount(firstGenInput int64, secondGenInput int64) int {

	criteria := func(restriction int64) criteriaFunc {
		return func(value int64) bool {
			return value%restriction == 0
		}
	}
	c1 := runGenerator(firstGenInput, genAFactor, criteria(genARestriction))
	c2 := runGenerator(secondGenInput, genBFactor, criteria(genBRestriction))
	counter := newJudge(extendedLimit, c1, c2)
	return counter
}
