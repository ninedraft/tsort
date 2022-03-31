package tsort_test

type graph [][]int

func (g graph) From(i int) []int { return g[i] }

func (g graph) All() []int {
	var nodes = make([]int, 0, len(g))
	for i := range g {
		nodes = append(nodes, i)
	}
	return nodes
}
