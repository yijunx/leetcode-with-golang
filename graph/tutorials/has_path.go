package main

import (
	"fmt"
	"slices"
)

func hasPathWithBFS(a map[string][]string, src string, dst string) bool {

	// lets use bfs
	// passedNodes := []string{}
	queue := []string{}

	for {
		// passedNodes = append(passedNodes, src)
		nextNodes, ok := a[src]
		if ok {
			if slices.Contains(nextNodes, dst) {
				return true
			}
			queue = append(queue, nextNodes...)
		}
		if len(queue) == 0 {
			break
		}
		// we can change the below line to make it DFS
		src, queue = queue[0], queue[1:]
	}
	return false
}

func hasPathWithDFSRecursive(a map[string][]string, src string, dst string) bool {

	for _, v := range a[src] {
		if slices.Contains(a[v], dst) {
			return true
		}
		hasPathWithDFSRecursive(a, v, dst)
	}
	return false

}

func main() {

	var adjacencyList map[string][]string = make(map[string][]string)

	adjacencyList["f"] = []string{"g", "i"}
	adjacencyList["g"] = []string{"h"}
	adjacencyList["h"] = []string{}
	adjacencyList["i"] = []string{"g", "k"}
	adjacencyList["j"] = []string{"i"}
	adjacencyList["k"] = []string{}

	// this is acyclic

	// f → g → h
	// ↓ ↗
	// i ← j
	// ↓
	// k
	fmt.Println(hasPathWithBFS(adjacencyList, "j", "f"))
	fmt.Println(hasPathWithBFS(adjacencyList, "f", "h"))

	fmt.Println(hasPathWithDFSRecursive(adjacencyList, "j", "f"))
	fmt.Println(hasPathWithDFSRecursive(adjacencyList, "f", "h"))

	// n = # nodes
	// e = # edges (e can be up to n^2, every possible pair of nodes has an edge!)
	// Time: O(e) coz we may travel through all edges
	// Space: O(n) size of our stack/queue
}
