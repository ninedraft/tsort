package tsort_test

import (
	"fmt"

	"github.com/ninedraft/tsort"
)

func ExampleSort() {
	var graph = [][]int{
		0: {},
		1: {0, 2},
		2: {0},
		3: {1, 2},
	}

	sorted, hasCycle := tsort.Sort(
		[]int{0, 1, 2, 3},
		func(i int) []int {
			return graph[i]
		})
	fmt.Println(sorted, hasCycle)
	// Output: [3 1 2 0] false
}

func ExampleIsSorted() {
	var graph = [][]int{
		0: {},
		1: {0, 2},
		2: {0},
		3: {1, 2},
	}

	isOrdered := tsort.IsSorted(
		[]int{0, 1, 2, 3},
		func(i int) []int {
			return graph[i]
		})
	fmt.Println(isOrdered)
	// Output: true
}
