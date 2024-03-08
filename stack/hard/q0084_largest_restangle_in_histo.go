package hard

import "fmt"

func largestRectangleArea(heights []int) int {

	n := len(heights)
	stack := []int{}
	lefts := make([]int, n)
	rights := make([]int, n)

	for i := 0; i < n; i++ {

		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			// pop all those bigger than me
			stack = stack[:len(stack)-1]
		}
		// pop untill all my left side is smaller or equal to me!

		if len(stack) == 0 {
			lefts[i] = 0
			// i can go left till the end
		} else {
			lefts[i] = stack[len(stack)-1] + 1
			// i can go left still i meet the previous one in the stack
		}

		stack = append(stack, i)
	}

	// pop all...
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			// pop all those bigger than me
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rights[i] = n - 1 // i can to rigt to the end
		} else {
			rights[i] = stack[len(stack)-1] - 1
			// i can only go right just before one prevoisone in the stack
		}

		stack = append(stack, i)
	}

	fmt.Println(lefts)
	fmt.Println(rights)

	maxArea := 0
	for i := 0; i < n; i++ {
		maxArea = max(maxArea, heights[i]*(rights[i]-lefts[i]+1))
	}

	return maxArea
}

func LargestRestangleInHisto() {

	heights := []int{2, 1, 5, 6, 2, 3}

	fmt.Println(largestRectangleArea(heights))

}
