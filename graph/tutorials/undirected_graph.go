package tutorials

import "fmt"

// func main() {

// }

// suppose we have an undirected graph where edges are

func hasPathChecksCycleBFS(a map[string][]string, src string, dst string) bool {
	queue := []string{}
	visitedNodes := make(map[string]bool)

	for {

		nextNodes, ok := a[src]

		if ok {
			for _, node := range nextNodes {
				if node == dst {
					return true
				}

				_, ok = visitedNodes[node]

				if !ok {
					// we have not visited
					visitedNodes[node] = true
					queue = append(queue, node)
				}
			}
			// queue = append(queue, nextNodes...)
		}

		if len(queue) == 0 {
			break
		}
		src, queue = queue[0], queue[1:]

	}

	return false
}

func hasPathChecksCycleDFSRecursive(a map[string][]string, src string, dst string, visitedNodes map[string]int) bool {
	fmt.Println("src:", src, "dst", dst)
	fmt.Println(src, "can go to", a[src])
	if src == dst {
		return true
	}

	fmt.Println("path from", src, "plus 1")
	visitedNodes[src] += 1

	for _, v := range a[src] {
		fmt.Println("we see", v)
		// make sure all path has been tried!!!!
		if !(visitedNodes[v] == len(a[v])) {
			return hasPathChecksCycleDFSRecursive(a, v, dst, visitedNodes)
		} else {
			fmt.Println("we exhausted the path from", v, "already!")
		}
	}
	return false
}

func hasPathChecksCycleDFS(a map[string][]string, src string, dst string) bool {
	visitedNodes := make(map[string]int)

	return hasPathChecksCycleDFSRecursive(
		a, src, dst, visitedNodes,
	)

}

func edgesToGraph(e [][2]string) map[string][]string {
	var adjacencyList map[string][]string = make(map[string][]string)

	for _, vs := range e {
		_, firstExists := adjacencyList[vs[0]]
		if firstExists {
			adjacencyList[vs[0]] = append(adjacencyList[vs[0]], vs[1])
		} else {
			adjacencyList[vs[0]] = []string{vs[1]}
		}
		_, secondExists := adjacencyList[vs[1]]
		if secondExists {
			adjacencyList[vs[1]] = append(adjacencyList[vs[1]], vs[0])
		} else {
			adjacencyList[vs[1]] = []string{vs[0]}
		}
	}

	return adjacencyList
}

func UndirectedCyclicGraphHasPath() {
	edges := [][2]string{
		{"i", "j"},
		{"k", "i"},
		{"k", "j"},
		{"m", "k"},
		{"k", "l"},
		{"o", "n"},
	}
	fmt.Println(edges)
	// first lets turn edges into a graph
	adjacencyList := edgesToGraph(edges)
	fmt.Println(adjacencyList)

	// we have a cycle i, j, k
	// thus we are easiy trapped in the infinite traversal

	// i ⎯ j
	// | ⟋
	// k ⎯ l
	// |
	// m

	// o ⎯ n

	// fmt.Println(hasPathChecksCycleBFS(adjacencyList, "i", "o"))
	// fmt.Println(hasPathChecksCycleDFS(adjacencyList, "i", "o"))

	// fmt.Println(hasPathChecksCycleBFS(adjacencyList, "i", "l"))
	// fmt.Println(hasPathChecksCycleDFS(adjacencyList, "i", "l"))

	// fmt.Println(hasPathChecksCycleBFS(adjacencyList, "m", "j"))
	// fmt.Println(hasPathChecksCycleDFS(adjacencyList, "m", "j"))

	fmt.Println(hasPathChecksCycleBFS(adjacencyList, "m", "o"))
	fmt.Println(hasPathChecksCycleDFS(adjacencyList, "m", "o"))
}
