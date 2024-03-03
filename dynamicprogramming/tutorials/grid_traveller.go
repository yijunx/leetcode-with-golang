package tutorials

import "fmt"

type size struct {
	n int
	m int
}

func gridTraveller(n int, m int) int {
	// well we know this for sure
	memo := map[size]int{
		{1, 1}: 1,
	}
	return gridTravellerWithMemo(n, m, memo)
}

func gridTravellerWithMemo(n int, m int, memo map[size]int) int {
	// check if outside
	if n == 0 {
		return 0
	}
	if m == 0 {
		return 0
	}
	// check if exists
	if steps, exists := memo[size{n, m}]; exists {
		return steps
	} else {
		// calculate new and save into memo
		newSteps := gridTravellerWithMemo(n-1, m, memo) + gridTravellerWithMemo(n, m-1, memo)
		memo[size{n, m}] = newSteps
		memo[size{m, n}] = newSteps
		return newSteps
	}
}

func gridTravellerWithTabu(n int, m int) int {
	grid := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		grid[i] = make([]int, m+1)
	}
	grid[1][1] = 1
	fmt.Println(grid)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j > 0 || i > 0 {
				grid[i+1][j+1] = grid[i][j+1] + grid[i+1][j]
			}
		}
	}
	return grid[n][m]
}

func GridTraveller() {

	// suppose we have a grid 2 * 3, n * m
	// how many ways from top left to bottom right
	// in a shortest fashion

	// * * *
	// * * *

	// equals to

	// * *
	// * *

	// plus

	// * * *

	// 1 1 1 1
	// 1 2 3 4
	// 1 3 6 10
	fmt.Println(gridTraveller(3, 3))
	fmt.Println(gridTravellerWithTabu(3, 3))
	fmt.Println(gridTraveller(4, 4))
	fmt.Println(gridTravellerWithTabu(4, 4))

}
