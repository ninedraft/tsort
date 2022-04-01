package tsort_test

import (
	"testing"

	"github.com/ninedraft/tsort"
)

func TestIsOrdered(test *testing.T) {
	t := func(name string, nodes []int, successors tsort.Successors[int], expected bool) {
		test.Run(name, func(test *testing.T) {
			var got = tsort.IsSorted(nodes, successors)
			if got != expected {
				switch {
				case expected:
					test.Errorf("nodes are expected to be sorted: %v", nodes)
				default:
					test.Errorf("nodes are expected to be not sorted: %v", nodes)
				}
			}
		})
	}

	t("empty nodes", []int{}, graph{
		0: {1},
		1: {},
	}.From, true)

	t("empty graph", []int{}, graph{}.From, true)

	t("sorted", []int{2, 1, 3, 4, 0}, graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}.From, true)

	t("sorted subgraph", []int{2, 1, 3}, graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}.From, true)

	t("not sorted", []int{1, 2, 3, 4}, graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}.From, true)

	t("not sorted subgraph", []int{1, 2, 3}, graph{
		0: {},
		1: {0},
		2: {0, 1},
		3: {2},
		4: {0, 1, 2, 3},
	}.From, true)
}
