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

func ConnectedComponentsCount() {
	graph := map[int][]int{
		3: {},
		4: {6},
		6: {4, 5, 7, 8},
		8: {6, 7},
		7: {6, 8},
		5: {6},
		1: {2},
		2: {1},
	}

	fmt.Println(graph)

	// 1 - 2
	//
	//     4
	//     |
	// 5 - 6 - 8
	//     | ⟋
	// 3   7

	// should return 3.. for above picture

	fmt.Println(connectedComponentsCountBFS(graph))
}
