package medium

import "fmt"

func longestPalindromicSubstringWithGrid(s string) string {
	// initialize the grid
	if len(s) <= 1 {
		return s
	}

	// maxLen := 1
	maxPal := s[0:1]
	grid := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		grid[i] = make([]bool, len(s))
	}

	for i := range s {
		grid[i][i] = true
		for j := 0; j < i; j++ {

		}
	}
	// babqw

	// T F F F F
	// F T F F F
	// T F T F F
	// F F F T F
	// F F F F T

	return maxPal
}

func longestPalindromicSubstringWithManacher(s string) string {
	return ""
}

func LongestPalindromicSubstring() {

	a := "babqwerrt"

	fmt.Println(longestPalindromicSubstringWithGrid(a))
	fmt.Println(longestPalindromicSubstringWithManacher(a))
}
