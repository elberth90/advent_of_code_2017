package tower

import (
	"strconv"
	"strings"
)

type node struct {
	children    []*node
	parent      *node
	name        string
	weight      int
	totalWeight int
}

func buildNode(input string) *node {
	data := strings.Split(input, " ")
	w, err := strconv.Atoi(strings.Trim(strings.Trim(data[1], "("), ")"))
	if err != nil {
		panic(err)
	}
	n := node{
		name:        data[0],
		weight:      w,
		children:    []*node{},
		totalWeight: w,
	}
	return &n
}

func connect(parent *node, child *node) {
	parent.children = append(parent.children, child)
	child.parent = parent
}

func buildTower(input string) map[string]*node {
	data := strings.Split(input, "\n")
	var nodeList = make(map[string]*node)
	for _, d := range data {
		n := buildNode(d)
		nodeList[n.name] = n

	}
	for _, d := range data {
		i := strings.Split(d, " ")
		if len(i) < 4 {
			continue
		}
		for _, j := range i[3:] {
			connect(nodeList[i[0]], nodeList[strings.Trim(j, ",")])
		}
	}

	for _, n := range nodeList {
		n.totalWeight = calculateTotal(n)
	}

	return nodeList
}

func calculateTotal(n *node) int {
	if len(n.children) == 0 {
		return n.weight
	}

	s := n.weight
	for _, c := range n.children {
		s += calculateTotal(c)
	}

	return s
}

// FindBottomProgram return name of the root program
func FindBottomProgram(input string) string {
	nodeList := buildTower(input)
	return getRoot(nodeList).name
}

func getRoot(nodeList map[string]*node) *node {
	for _, n := range nodeList {
		if n.parent == nil {
			return n
		}
	}
	return nil
}

// GetWeightToBalance return proper weight of node to make sure that tree is balanced
func GetWeightToBalance(input string) int {
	nodeList := buildTower(input)

	var fixedWeight int
	for n := getRoot(nodeList); n != nil && len(n.children) > 2; {
		w := n.children[0].totalWeight
		w2 := n.children[1].totalWeight
		if n.children[2].totalWeight == w2 {
			w = w2
		}
		var next *node
		for _, n := range n.children {
			if n.totalWeight != w {
				next = n
				fixedWeight = n.weight - (n.totalWeight - w)
			}
		}
		n = next
	}

	return fixedWeight
}
