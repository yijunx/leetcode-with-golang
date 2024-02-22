package tutorials

import (
	"fmt"
)

func islandCount(g [][]string) int {

	type location struct {
		i int
		j int
	}

	getNeighborLocs := func(loc location, gridSize int) []location {
		result := []location{}

		if loc.i-1 >= 0 {
			result = append(result, location{loc.i - 1, loc.j})
		}

		if loc.j-1 >= 0 {
			result = append(result, location{loc.i, loc.j - 1})
		}

		if loc.i+1 < gridSize-1 {
			result = append(result, location{loc.i + 1, loc.j})
		}

		if loc.j+1 < gridSize-1 {
			result = append(result, location{loc.i, loc.j + 1})
		}

		return result
	}

	visitedNodes := map[location]bool{}
	landingLocs := []location{}
	gridSize := len(g)
	count := 0

	for i, row := range g {
		for j, value := range row {

			// only start if the current location is not checked

			currentLoc := location{i: i, j: j}

			if value != "L" {
				// we only focus on land
				// if water we just skip
				continue
			}

			if _, exists := visitedNodes[currentLoc]; exists {
				continue
			}

			landingLocs = append(landingLocs, currentLoc)

			// for lets start bfs
			queue := []location{}

			for {
				// now we see a land mass, time to go over all
				// the cells possible of this land mass
				fmt.Println("we are at", currentLoc)
				visitedNodes[currentLoc] = true

				for _, neighborLoc := range getNeighborLocs(currentLoc, gridSize) {
					fmt.Println("neighbors are", neighborLoc)
					if _, exists := visitedNodes[neighborLoc]; exists {
						// well we have been here...
					} else {
						// add to queue to neighborLoc is landmass also...
						if g[neighborLoc.i][neighborLoc.j] == "L" {
							queue = append(queue, neighborLoc)
							visitedNodes[neighborLoc] = true
						}
					}
				}

				if len(queue) == 0 {
					count += 1
					break
				}

				currentLoc, queue = queue[0], queue[1:]
			}
		}
	}
	fmt.Println(landingLocs)
	fmt.Println(visitedNodes)
	fmt.Println(count)
	return len(landingLocs)
}

func IslandCount() {

	grid := [][]string{
		{"W", "L", "W", "W", "L", "W"},
		{"L", "L", "W", "W", "L", "W"},
		{"W", "L", "W", "W", "w", "W"},
		{"W", "w", "W", "L", "L", "W"},
		{"W", "L", "W", "L", "L", "W"},
		{"W", "W", "W", "W", "W", "W"},
	}

	fmt.Println(islandCount(grid))
}
