package gpu

import (
	"math"
	"reflect"
	"strings"
)

// ClosestParticle return id of the particle that is closest to the center
func ClosestParticle(input string) int {
	particleList := loadParticles(input)
	distance := map[int]int{}
	last := -1
	counter := 0

	for {
		for _, p := range particleList {
			p.Tick()
			distance[p.id] = p.Distance()
		}
		id := getClosestParticleID(distance)
		if last == id {
			counter++
		} else {
			counter = 0
		}
		last = id
		if counter >= 100000 {
			break
		}
	}
	return last
}

// AfterCollisions return amount of particles after collisions
func AfterCollisions(input string) int {
	particleList := loadParticles(input)
	counter := 0
	for {
		for _, p := range particleList {
			p.Tick()
		}

		newP, ok := calculateCollisions(particleList)
		if !ok {
			counter++
		} else {
			counter = 0
		}
		if counter >= 100 {
			break
		}
		particleList = newP
	}
	return len(particleList)
}

func calculateCollisions(particles []*particle) ([]*particle, bool) {
	toRemove := []int{}
	for _, originalP := range particles {
		for _, toCompareP := range particles {
			if originalP.id == toCompareP.id {
				continue
			}

			if reflect.DeepEqual(originalP.position, toCompareP.position) {
				toRemove = append(toRemove, originalP.id, toCompareP.id)
			}
		}
	}

	newP := []*particle{}
	if len(toRemove) > 0 {
		for _, p := range particles {
			add := true
			for _, rP := range toRemove {
				if rP == p.id {
					add = false
					break
				}
			}
			if add {
				newP = append(newP, p)
			}
		}
	} else {
		newP = particles
	}

	return newP, len(newP) != len(particles)
}

func getClosestParticleID(distance map[int]int) int {
	lowestID := 0
	lowest := int(math.Abs(0 - float64(distance[lowestID])))

	for id, d := range distance {
		diff := int(math.Abs(0 - float64(d)))
		if diff < lowest {
			lowest = diff
			lowestID = id
		}
	}

	return lowestID
}

func loadParticles(input string) []*particle {
	rawParticleList := strings.Split(strings.Trim(input, "\n"), "\n")
	var particleList = make([]*particle, len(rawParticleList))

	for i, p := range rawParticleList {
		p = strings.TrimSpace(p)
		particleList[i] = fromRawData(i, p)
	}

	return particleList
}
