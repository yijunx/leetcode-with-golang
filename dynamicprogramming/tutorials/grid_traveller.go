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
		return newSteps
	}
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

	// 0 1 1 1
	// 1 2 3 4
	// 1 3 6 10
	fmt.Println(gridTraveller(50, 50))

}
