[![Go Reference](https://pkg.go.dev/badge/github.com/ninedraft/tsort.svg)](https://pkg.go.dev/github.com/ninedraft/tsort) [![](https://goreportcard.com/badge/github.com/ninedraft/tsort)](https://goreportcard.com/report/github.com/ninedraft/tsort) [![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) [![Github tag](https://badgen.net/github/tag/ninedraft/tsort)](https://github.com/ninedraft/tsort/tags/)
# tsort

## Installation

```
go get github.com/ninedraft/tsort
```

## Description
A small [topological ordering/sorting](https://en.wikipedia.org/wiki/Topological_sorting) generic library.

It doesn't define any graph format, instead you can just pass a single callback.

```go
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
```

It also provides a IsSorted function, which checks if provided nodes are sorted on graph:

```go
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

```

## Deps

Zero dependencies (beside of std and x/exp).
