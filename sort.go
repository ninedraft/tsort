package tsort

// Sort performs topological sort of nodes on provided directed graph.
// Provided nodes are sorted in place.
// Returns true, if provided graph cycle is detected.
func Sort[S ~[]N, N comparable](nodes S, successors Successors[N]) (hasCycle bool) {
	var s = sorter[N]{
		colored:    make(map[N]nodeColor, len(nodes)),
		sorted:     nodes,
		successors: successors,
		offset:     len(nodes) - 1,
	}
	for _, node := range nodes {
		s.colored[node] = colorWhite
	}
	hasCycle = s.sort(nodes)
	return hasCycle
}

type nodeColor int8

const (
	colorWhite nodeColor = iota
	colorBlack
	colorGray
)

type sorter[N comparable] struct {
	offset     int
	sorted     []N
	successors Successors[N]
	colored    map[N]nodeColor
}

func (s *sorter[N]) sort(nodes []N) (hasCycle bool) {
	// using Robert Tarjan algorythm
	for i := len(nodes) - 1; i >= 0; i-- {
		var node = nodes[i]
		var color, isKnown = s.colored[node]
		// ignoring unmentioned parts of graph
		// we are sorting only nodes slice
		if !isKnown {
			continue
		}
		switch color {
		case colorBlack:
			continue
		case colorGray:
			return true
		case colorWhite:
			s.colored[node] = colorGray
			var sc = s.successors(node)
			var hasCycle = s.sort(sc)
			if hasCycle {
				return hasCycle
			}
			s.colored[node] = colorBlack
			s.sorted[s.offset] = node
			s.offset--
		}
	}
	return false
}
