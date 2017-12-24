package gpu

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const r = `p=<(?P<position>\s?-?\d+,\s?-?\d+,\s?-?\d+)>,\sv=<(?P<velocity>\s?-?\d+,\s?-?\d+,\s?-?\d+)>,\sa=<(?P<acceleration>\s?-?\d+,\s?-?\d+,\s?-?\d+)>`

type particle struct {
	id           int
	position     coord
	velocity     coord
	acceleration coord
}

func (p *particle) Tick() {
	p.velocity.x += p.acceleration.x
	p.velocity.y += p.acceleration.y
	p.velocity.z += p.acceleration.z

	p.position.x += p.velocity.x
	p.position.y += p.velocity.y
	p.position.z += p.velocity.z
}

func (p *particle) Distance() int {
	return int(math.Abs(float64(p.position.x)) + math.Abs(float64(p.position.y)) + math.Abs(float64(p.position.z)))
}

func fromRawData(id int, input string) *particle {
	params := splitParams(input)
	return &particle{
		id:           id,
		position:     paramToCoord(params["position"]),
		velocity:     paramToCoord(params["velocity"]),
		acceleration: paramToCoord(params["acceleration"]),
	}
}

func paramToCoord(param string) coord {
	p := strings.Split(param, ",")

	x, err := strconv.Atoi(strings.TrimSpace(p[0]))
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(strings.TrimSpace(p[1]))
	if err != nil {
		panic(err)
	}

	z, err := strconv.Atoi(strings.TrimSpace(p[2]))
	if err != nil {
		panic(err)
	}

	return coord{x: x, y: y, z: z}
}

func splitParams(input string) map[string]string {
	r := regexp.MustCompile(r)
	match := r.FindStringSubmatch(input)
	paramsMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

type coord struct {
	x int
	y int
	z int
}
