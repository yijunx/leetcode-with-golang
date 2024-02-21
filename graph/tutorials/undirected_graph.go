package tutorials

import "fmt"

// func main() {

// }

// suppose we have an undirected graph where edges are

func hasPathChecksCycleBFS(a map[string][]string, src string, dst string) bool {
	queue := []string{}
	startedFromNodes := make(map[string]bool)

	for {

		nextNodes, ok := a[src]

		if ok {
			for _, node := range nextNodes {
				if node == dst {
					return true
				}

				_, ok = startedFromNodes[node]

				if !ok {
					// we have not visited
					startedFromNodes[node] = true
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
	// well here golang is so cool, it starts from zero
	// no need to check if src is in the visited nodes!
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

func hasPathChecksCycleDFSRecursive2(a map[string][]string, src string, dst string, visitedNodes map[string]bool) bool {
	fmt.Println("src:", src, "dst", dst)
	fmt.Println(src, "can go to", a[src])
	if src == dst {
		return true
	}

	// check if visited
	if _, exists := visitedNodes[src]; exists {
		return false
	}
	// well here golang is so cool, it starts from zero
	// no need to check if src is in the visited nodes!
	visitedNodes[src] = true

	for _, v := range a[src] {
		fmt.Println("we see", v)

		if hasPathChecksCycleDFSRecursive2(a, v, dst, visitedNodes) {
			return true
		}
		// here we cannot do
		// return hasPathChecksCycleDFSRecursive(a, v, dst, visitedNodes)
		// because we want to keep the search
		// if we want to do this, needs to check all paths checked like
		// hasPathChecksCycleDFSRecursive

	}
	return false
}

func hasPathChecksCycleDFS(a map[string][]string, src string, dst string) bool {
	visitedNodes := make(map[string]int)

	return hasPathChecksCycleDFSRecursive(
		a, src, dst, visitedNodes,
	)

}

func hasPathChecksCycleDFS2(a map[string][]string, src string, dst string) bool {
	visitedNodes := make(map[string]bool)

	return hasPathChecksCycleDFSRecursive2(
		a, src, dst, visitedNodes,
	)

}

func edgesToGraph(e [][2]string) map[string][]string {
	var adjacencyList map[string][]string = make(map[string][]string)

	for _, vs := range e {
		a, b := vs[0], vs[1]
		adjacencyList[a] = append(adjacencyList[a], b)
		adjacencyList[b] = append(adjacencyList[b], a)

		// below code is too python..
		// _, firstExists := adjacencyList[a]
		// if firstExists {
		// 	adjacencyList[a] = append(adjacencyList[a], b)
		// } else {
		// 	adjacencyList[a] = []string{b}
		// }
		// _, secondExists := adjacencyList[b]
		// if secondExists {
		// 	adjacencyList[b] = append(adjacencyList[b], a)
		// } else {
		// 	adjacencyList[b] = []string{a}
		// }
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
	fmt.Println(hasPathChecksCycleDFS2(adjacencyList, "i", "l"))
}
