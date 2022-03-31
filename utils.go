package tsort

type set[N comparable] map[N]struct{}

func newSet[N comparable](items []N) set[N] {
	var s = make(set[N], len(items))
	for _, item := range items {
		s.add(item)
	}
	return s
}

func (s set[N]) add(node N) {
	s[node] = struct{}{}
}

func (s set[N]) contains(node N) bool {
	var _, ok = s[node]
	return ok
}

func (s set[N]) containsAny(nodes []N) bool {
	for _, node := range nodes {
		if s.contains(node) {
			return true
		}
	}
	return false
}
