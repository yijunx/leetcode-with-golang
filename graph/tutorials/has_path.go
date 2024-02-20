package tutorials

import "fmt"

// n = # nodes
// e = # edges (e can be up to n^2, every possible pair of nodes has an edge!)
// Time: O(e) coz we may travel through all edges
// Space: O(n) size of our stack/queue

func HasPathWithBFS(a map[string][]string, src string, dst string) bool {

	// lets use bfs
	// passedNodes := []string{}
	queue := []string{}

	for {
		// passedNodes = append(passedNodes, src)
		nextNodes := a[src]
		for _, v := range nextNodes {
			if v == dst {
				return true
			}
			queue = append(queue, v)
		}
		if len(queue) == 0 {
			break
		}
		// we can change the below line to make it DFS
		src, queue = queue[0], queue[1:]
	}
	return false
}

func HasPathWithDFSRecursive(a map[string][]string, src string, dst string) bool {

	if src == dst {
		return true
	}
	for _, v := range a[src] {
		return HasPathWithDFSRecursive(a, v, dst)
	}
	return false

}

func HasPath() {
	// get to called by the main.go
	var adjacencyList map[string][]string = make(map[string][]string)

	adjacencyList["f"] = []string{"g", "i"}
	adjacencyList["g"] = []string{"h"}
	adjacencyList["h"] = []string{}
	adjacencyList["i"] = []string{"g", "k"}
	adjacencyList["j"] = []string{"i"}
	adjacencyList["k"] = []string{}
	// f → g → h
	// ↓ ↗
	// i ← j
	// ↓
	fmt.Println(HasPathWithBFS(adjacencyList, "j", "f"))
	fmt.Println(HasPathWithBFS(adjacencyList, "f", "h"))

	fmt.Println(HasPathWithDFSRecursive(adjacencyList, "j", "f"))
	fmt.Println(HasPathWithDFSRecursive(adjacencyList, "f", "h"))
}
