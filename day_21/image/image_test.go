package image

import "testing"

func TestMakeArt(t *testing.T) {
	useCases := map[string]struct {
		rules      string
		iterations int
		onPixels   int
	}{
		"a": {
			rules:      "../.# => ##./#../...\n.#./..#/### => #..#/..../..../#..#",
			iterations: 2,
			onPixels:   12,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			onPixels := MakeArt(uc.rules, uc.iterations)
			if onPixels != uc.onPixels {
				t.Fatalf("Expected number of on pixes was `%d`, got `%d`", uc.onPixels, onPixels)
			}
		})
	}
}
