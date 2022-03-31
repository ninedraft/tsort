package tsort

// Successors describes a generic function, which resolves successors of node.
// Successors of node are nodes such that there exists directed edges from node to each successor.
type Successors[N comparable] func(node N) []N
