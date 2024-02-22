package tutorials

import "fmt"

func largestComponentBFS(g map[int][]int) int {
	arrivedAtNodes := map[int]bool{}
	startedFromNodes := map[int]bool{}
	maxSize := 0

	for startingNode := range g {
		_, exists := arrivedAtNodes[startingNode]

		if exists {
			fmt.Println("we have covered", startingNode)
			continue
		} else {
			fmt.Println("doing", startingNode)
		}
		// now we start to travers with BFS
		// and BFS requires a queue and starting loc
		queue := []int{}
		src := startingNode
		traveredNodesThisTime := map[int]bool{}
		for {
			nextNodes := g[src]
			traveredNodesThisTime[src] = true
			arrivedAtNodes[src] = true

			// add stuff to queue
			for _, nextNode := range nextNodes {
				// check if we have started from next node before
				_, exists := startedFromNodes[nextNode]

				if !exists {
					queue = append(queue, nextNode)
					startedFromNodes[src] = true
				}
			}
			if len(queue) == 0 {
				break
			}
			src, queue = queue[0], queue[1:]
		}
		fmt.Println("size of this cluster is", len(traveredNodesThisTime))

		maxSize = max(maxSize, len(traveredNodesThisTime))
	}
	return maxSize
}

func largestComponentDFS(g map[int][]int) int {
	fmt.Println("this is dfs")
	visitedNodes := map[int]bool{}
	traversedNodesCountFromStartingNode := map[int]int{}
	maxSize := 0
	for startingNode := range g {
		if _, exists := visitedNodes[startingNode]; exists {
			continue
		}
		fmt.Println(visitedNodes)
		largestComponentDFSRecursive(g, startingNode, visitedNodes, startingNode, traversedNodesCountFromStartingNode)
	}
	fmt.Println(traversedNodesCountFromStartingNode)
	for _, v := range traversedNodesCountFromStartingNode {
		maxSize = max(v, maxSize)
	}
	return maxSize
}

func largestComponentDFSRecursive(g map[int][]int, src int, visitedNodes map[int]bool, startingNode int, traversedNodesCountFromStartingNode map[int]int) bool {
	_, exists := visitedNodes[src]
	if exists {
		fmt.Println("well we visited", src)
		return false
	}
	visitedNodes[src] = true
	traversedNodesCountFromStartingNode[startingNode] += 1

	for _, neighbor := range g[src] {
		if largestComponentDFSRecursive(g, neighbor, visitedNodes, startingNode, traversedNodesCountFromStartingNode) {
			return true
		}
	}
	return false
}

func LargestComponent() {

	adjacencyList := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0, 8},
		8: {0, 4},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}

	//     5
	//     | ⟍
	// 1 - 0 - 8  size=4
	//
	// 4 - 2      size=3
	// | ⟋
	// 3

	// thus we should return 4

	fmt.Println(adjacencyList)
	fmt.Println(largestComponentBFS(adjacencyList))
	fmt.Println(largestComponentDFS(adjacencyList))

}
