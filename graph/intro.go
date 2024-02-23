package graph

// graph is nodes + edges
// directed graph, edge has directions
// undirected graph
// neighbor is what is accessable from current node
// we can use a hashmap to store the adjacency

import (
	"fmt"
)

func printGraph(a map[string][]string) {
	for k, v := range a {
		for _, item := range v {
			fmt.Printf("%v -> %v\n", k, item)
		}
	}
}

func dfs(a map[string][]string, startingLoc string) {
	// going deeper first!!!
	// a b d ...
	// drilling down...
	// using stack, add to top, remove from top
	fmt.Println("begin dfs!!")

	result := []string{}
	stack := []string{}

	for {
		fmt.Println("appending", startingLoc)
		nextLocs, ok := a[startingLoc]
		fmt.Println("current starting loc", startingLoc, "goes to", nextLocs)
		result = append(result, startingLoc)
		if ok {
			stack = append(stack, nextLocs...)
		}
		fmt.Println("stack is ", stack)
		if len(stack) == 0 {
			break
		}
		// POP last
		startingLoc = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	fmt.Println(result)
}

func dfsRecursive(a map[string][]string, startingLoc string) {
	// here we are LITERALLY using a stack
	// the call stack
	// thats why bfs cannot go with this technique
	// as it requires queue
	fmt.Println(startingLoc)
	for _, v := range a[startingLoc] {
		dfsRecursive(a, v)
	}
}

func bfs(a map[string][]string, startingLoc string) {
	// going wider first!!!
	// a b c ...
	// tend to explore nearby evenly
	// using a queue, add to back, remove from front
	fmt.Println("begin bfs!!")

	result := []string{}
	queue := []string{}

	for {
		fmt.Println("appending", startingLoc)
		nextLocs, ok := a[startingLoc]
		fmt.Println("current starting loc", startingLoc, "goes to", nextLocs)
		result = append(result, startingLoc)
		if ok {
			// there are destinations
			queue = append(queue, nextLocs...)
		}
		fmt.Println("queue is ", queue)
		if len(queue) == 0 {
			break
		}
		// POP first
		startingLoc = queue[0]
		queue = queue[1:]
	}

	fmt.Println(result)
}

func GraphIntro() {

	var adjacencyList map[string][]string = make(map[string][]string)

	adjacencyList["a"] = []string{"b", "c"}
	adjacencyList["b"] = []string{"d"}
	adjacencyList["c"] = []string{"e"}
	adjacencyList["d"] = []string{"f"}
	adjacencyList["e"] = []string{}
	adjacencyList["f"] = []string{}
	printGraph(adjacencyList)

	// well lets draw this
	//        a
	//      b   c
	//    d       e
	//  f

	dfs(adjacencyList, "a")
	dfsRecursive(adjacencyList, "a")
	bfs(adjacencyList, "a")
}
