package gpu

import "testing"

func TestClosestParticle(t *testing.T) {
	useCases := map[string]struct {
		input           string
		closestParticle int
	}{
		"": {
			input:           "p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>\np=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>",
			closestParticle: 0,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			particle := ClosestParticle(uc.input)
			if particle != uc.closestParticle {
				t.Fatalf("Expected particle was `%d`, got `%d`", uc.closestParticle, particle)
			}

		})
	}
}

func TestAfterCollisions(t *testing.T) {
	useCases := map[string]struct {
		input           string
		closestParticle int
	}{
		"": {
			input:           "p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>\np=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>\np=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>\np=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>",
			closestParticle: 1,
		},
	}

	for ucName, uc := range useCases {
		t.Run(ucName, func(t *testing.T) {
			particle := AfterCollisions(uc.input)
			if particle != uc.closestParticle {
				t.Fatalf("Expected particle was `%d`, got `%d`", uc.closestParticle, particle)
			}

		})
	}
}
