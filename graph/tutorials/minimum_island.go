package tutorials

import "fmt"

type loc struct {
	i int
	j int
}

func minimumIslandBFS(g [][]string) int {

	visitedLocs := map[loc]bool{}
	islandMap := map[loc][]loc{}
	n, m := len(g), len(g[0])

	for i, row := range g {
		for j, value := range row {
			if value == "W" {
				continue
			}
			landingLoc := loc{i: i, j: j}
			src := loc{i: i, j: j}
			if _, exists := visitedLocs[src]; exists {
				continue
			}

			// here we can start our bfs
			queue := []loc{}

			fmt.Println("src is", src)

			islandMap[landingLoc] = append(islandMap[landingLoc], landingLoc)
			for {

				visitedLocs[src] = true

				// add neighbor into the queue
				neighbor_locs := []loc{
					{i: src.i, j: src.j + 1},
					{i: src.i, j: src.j - 1},
					{i: src.i + 1, j: src.j},
					{i: src.i - 1, j: src.j},
				}
				for _, neighbor := range neighbor_locs {
					if neighbor.i < 0 || neighbor.i == n {
						continue
					}

					if neighbor.j < 0 || neighbor.j == m {
						continue
					}

					// now the location is good, check if its land
					if g[neighbor.i][neighbor.j] == "L" {
						fmt.Println("checking on neighbor", neighbor)
						if _, exists := visitedLocs[neighbor]; !exists {
							visitedLocs[neighbor] = true
							fmt.Println("adding", neighbor, "to", landingLoc)
							islandMap[landingLoc] = append(islandMap[landingLoc], neighbor)
							queue = append(queue, neighbor)
						}
					}
				}
				if len(queue) == 0 {
					break
				}

				src, queue = queue[0], queue[1:]
			}
		}
	}

	minSize := n * m
	fmt.Println(islandMap)

	for _, locs := range islandMap {
		minSize = min(minSize, len(locs))
	}
	return minSize
}

func minimumIslandDFS(g [][]string) int {

	size := exploreGrid(g)
	return size
}

func exploreGrid(g [][]string) int {
	return 0
}

func MinimumIsland() {
	grid := [][]string{
		{"W", "L", "W", "W", "L", "W"},
		{"L", "L", "W", "W", "L", "W"},
		{"W", "L", "W", "W", "W", "W"},
		{"W", "W", "W", "L", "L", "W"},
		{"W", "L", "W", "L", "L", "W"},
		{"W", "W", "W", "W", "W", "W"},
	}

	fmt.Println(minimumIslandBFS(grid))
	fmt.Println("here comes dfs")
	fmt.Println(minimumIslandDFS(grid))
}
