package medium

import "fmt"

// func printRuneSlice(rs []rune) {
// 	result := ""

// 	for _, r := range rs {
// 		result += string(r)
// 	}
// 	fmt.Println(result, "length is", len(rs))
// }

func removeKdigits(num string, k int) string {
	stack := []rune{}

	// remove the most significant big numbers
	// 1, k = 3, we see 4, 4 gets into stack
	// 1 4, k = 3, we see 3, 4 > 3, 4 gets popped!!, 3 added
	// 1 3, k = 2, we see 2, 3 > 2, 3 gets popped!!, 2 added
	// 1 2, k = 1, we see 2, 2 gets into stack
	// 1 2 2, k = 1, we see 1, 2 > 1, 2 gets popped!!, 1 added
	// 1 2 1, k = 0, just append
	// 1 2 1 9

	// 1 2 3 0, and k = 3
	// 1 2 3 , then see 0,
	// all 1 2 3 gets popped..
	// lets with nothing
	for _, digit := range num {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > digit {
			// pop from stack
			// the incoming number is smaller than the top of the stack
			// we cannot use it,
			stack = stack[:len(stack)-1]
			// used one chance
			k -= 1
		}

		if len(stack) > 0 || digit != '0' {
			// only skip the starting zeros
			stack = append(stack, digit)
		}
	}

	if k > 0 {
		// notice the max here
		stack = stack[:max(len(stack)-k, 0)]
	}

	result := ""

	for _, r := range stack {
		result += string(r)
	}
	// fmt.Println(result)
	if result == "" {
		return "0"
	}
	return result
}

func RemoveKDigits() {

	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10001", 4))
}
