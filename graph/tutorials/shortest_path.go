package tutorials

import "fmt"

func transformEdgesToGraph(edges [][2]string) map[string][]string {
	graph := map[string][]string{}

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	return graph
}

func shortestPath(g map[string][]string, src string, dst string) int {
	// return the number of edges
	// i guess bfs should be the way to go..
	queue := []string{}
	visitedNodes := map[string]bool{}
	minStepsToDst := map[string]int{src: 0}

	for {

		visitedNodes[src] = true
		neighbors := g[src]

		fmt.Println("we are at", src)
		// fmt.Println("queue is", queue)
		// fmt.Println("neighbors are", neighbors)

		for _, neighbor := range neighbors {

			if _, exists := minStepsToDst[neighbor]; !exists {
				minStepsToDst[neighbor] = minStepsToDst[src] + 1
			}

			if neighbor == dst {
				return minStepsToDst[neighbor]
			}
			if _, exists := visitedNodes[neighbor]; exists {
				// means we have been here..
				// fmt.Println("we have visited", neighbor)
			} else {
				// fmt.Println("appending", neighbor)
				queue = append(queue, neighbor)
			}
		}
		if len(queue) == 0 {
			break
		}
		src, queue = queue[0], queue[1:]

	}
	return 0
}

func ShortestPath() {

	edges := [][2]string{
		{"w", "x"},
		{"y", "x"},
		{"y", "z"},
		{"z", "v"},
		{"w", "v"},
	}

	fmt.Println(edges)
	fmt.Println(transformEdgesToGraph(edges))

	//
	// x - y - z
	// |       |
	// w ----- v

	fmt.Println(shortestPath(transformEdgesToGraph(edges), "x", "v"))
}
