package coprocessor

import "testing"

func TestHowManyMul(t *testing.T) {
	useCases := map[string]struct {
		instructions string
		mulCounter   int
	}{
		"aoc real data": {
			instructions: "set b 79\nset c b\njnz a 2\njnz 1 5\nmul b 100\nsub b -100000\nset c b\nsub c -17000\nset f 1\nset d 2\nset e 2\nset g d\nmul g e\nsub g b\njnz g 2\nset f 0\nsub e -1\nset g e\nsub g b\njnz g -8\nsub d -1\nset g d\nsub g b\njnz g -13\njnz f 2\nsub h -1\nset g b\nsub g c\njnz g 2\njnz 1 3\nsub b -17\njnz 1 -23",
			mulCounter:   5929,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			mulCounter := HowManyMul(uc.instructions)
			if mulCounter != uc.mulCounter {
				t.Fatalf("Expected mul counter was `%d`, got `%d`", uc.mulCounter, mulCounter)
			}
		})
	}
}
