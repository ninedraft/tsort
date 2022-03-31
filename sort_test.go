package tsort_test

import (
	"encoding/binary"
	"testing"
	"tsort"
)

func TestSort(test *testing.T) {
	var g = graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}

	var nodes = []int{2, 1, 3, 4, 0}
	var hasCycle = tsort.Sort(nodes, g.From)
	if hasCycle {
		test.Fatal("graph has cycle")
	}

	assertOrdered(test, nodes, g.From)
}

func TestSort_Empty(test *testing.T) {
	var g = graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}

	var nodes = []int{}
	var hasCycle = tsort.Sort(nodes, g.From)
	if hasCycle {
		test.Fatal("graph has cycle")
	}

	assertOrdered(test, nodes, g.From)
}

func TestSort_Cycled(test *testing.T) {
	var g = graph{
		0: {1},
		1: {2},
		2: {3},
		3: {4},
		4: {0},
	}
	var nodes = g.All()
	var hasCycle = tsort.Sort(nodes, g.From)
	if !hasCycle {
		test.Fatalf("graph expected to be cycled: %v", nodes)
	}
}

func FuzzSort(fuzz *testing.F) {
	var g = graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}
	fuzz.Fuzz(func(test *testing.T, data []byte) {
		var nodes = decodeInts(data)
		for _, node := range nodes {
			if node < 0 || node > 4 {
				return
			}
		}
		test.Log("nodes", nodes)
		if !tsort.Sort(nodes, g.From) {
			assertOrdered(test, nodes, g.From)
		}
	})
}

func decodeInts(data []byte) []int {
	size := len(data) / 8
	offset := 0
	var ints = make([]int, 0, size)
	for i := 0; i < size; i++ {
		x, n := binary.Varint(data[offset:])
		if n <= 0 {
			return ints
		}
		ints = append(ints, int(x))
	}
	return ints
}

func assertOrdered[N comparable](t testing.TB, nodes []N, successors tsort.Successors[N]) {
	t.Helper()
	var seen = set[N]{}
	for i, node := range nodes {
		seen.add(node)
		sc := successors(node)
		if seen.containsAny(sc) {
			t.Errorf("nodes %v: node %v is out of order after %v because contains %v",
				nodes, node, nodes[:i], sc)
			return
		}
	}
}

type set[N comparable] map[N]struct{}

func (s set[N]) add(node N) {
	s[node] = struct{}{}
}

func (s set[N]) containsAny(nodes []N) bool {
	for _, node := range nodes {
		var _, ok = s[node]
		if ok {
			return true
		}
	}
	return false
}
