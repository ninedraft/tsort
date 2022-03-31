package tsort

// IsOrdered returns true, if nodes of graph are topologically sorted.
func IsOrdered[S ~[]N, N comparable](nodes S, successors Successors[N]) bool {
	var seen = set[N]{}
	for _, node := range nodes {
		seen.add(node)
		sc := successors(node)
		if seen.containsAny(sc) {
			return true
		}
	}
	return len(nodes) == 0
}
