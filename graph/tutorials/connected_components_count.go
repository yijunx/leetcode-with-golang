package tutorials

import "fmt"

func connectedComponentsCountBFS(g map[int][]int) int {
	count := 0
	startedFromNodes := make(map[int]bool)
	arrivedAtNodes := make(map[int]bool)

	for currentNode := range g {
		_, exists := arrivedAtNodes[currentNode]
		if exists {
			fmt.Println(currentNode, "is visited")
			continue
		}
		src := currentNode
		queue := []int{}

		for {
			nextNodes, ok := g[src]
			if ok {
				for _, v := range nextNodes {
					arrivedAtNodes[v] = true
					// below code is to avoid cycle
					_, exists := startedFromNodes[src]
					if !exists {
						queue = append(queue, v)
						startedFromNodes[src] = true
					}
				}
			}

			if len(queue) == 0 {
				count += 1
				break
			}
			// this is the bfs
			src, queue = queue[0], queue[1:]
		}
	}
	return count
}

func connectedComponentsCountDFSRecursive(g map[int][]int) int {
	nodesTraversedPathsCount := make(map[int]int)
	arrivedAtNodes := make(map[int]bool)
	count := 0
	for currentNode := range g {
		_, exists := arrivedAtNodes[currentNode]
		if !exists {
			count += 1
		}
		dfs(g, currentNode, arrivedAtNodes, nodesTraversedPathsCount)
	}
	return count
}

func dfs(g map[int][]int, src int, arrivedAtNodes map[int]bool, nodesTraversedPathsCount map[int]int) {
	arrivedAtNodes[src] = true
	if nodesTraversedPathsCount[src] == len(g[src]) {
		fmt.Println("we have walked all paths from", src)
		return
	}
	nodesTraversedPathsCount[src] += 1
	for _, neighbor := range g[src] {
		// arrivedAtNodes[neighbor] = true
		// if nodesTraversedPathsCount[neighbor] == len(g[neighbor]) {
		// 	fmt.Println("we have walked all paths from", neighbor)
		// 	break
		// } else {
		dfs(g, neighbor, arrivedAtNodes, nodesTraversedPathsCount)
		//}
	}
}

func connectedComponentsCountDFSRecursive2(g map[int][]int) int {
	checkedNodes := make(map[int]bool)
	count := 0
	for currentNode := range g {
		if dfs2(g, currentNode, checkedNodes) {
			count += 1
		}
	}
	return count
}

func dfs2(g map[int][]int, src int, checkedNodes map[int]bool) bool {
	_, exists := checkedNodes[src]
	if exists {
		return false
	} else {
		checkedNodes[src] = true
	}
	for _, neighbor := range g[src] {
		dfs2(g, neighbor, checkedNodes)
	}
	return true
}

func ConnectedComponentsCount() {
	graph := map[int][]int{
		3: {},
		4: {6, 5},
		6: {4, 5, 7, 8},
		8: {6, 7},
		7: {6, 8},
		5: {6, 4},
		1: {2},
		2: {1},
	}

	fmt.Println(graph)

	// 1 - 2
	//
	//     4
	//   ⟋ |
	// 5 - 6 - 8
	//     | ⟋
	// 3   7

	// should return 3.. for above picture

	fmt.Println(connectedComponentsCountBFS(graph))
	fmt.Println(connectedComponentsCountDFSRecursive(graph))
	fmt.Println(connectedComponentsCountDFSRecursive2(graph))
}
