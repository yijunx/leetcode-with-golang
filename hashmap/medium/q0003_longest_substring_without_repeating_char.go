package main

// Example 1:

// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

import "fmt"

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	var appearance map[int]int = make(map[int]int)
	// maps value to location
	var currentRecord int = 0
	var startingLoc int = 0
	var maxRecord int = 0
	for i, v := range s {
		fmt.Printf("We have %v\n", v)
		var lastAppear, found = appearance[int(v)]

		if found && lastAppear >= startingLoc {
			fmt.Printf("duplication found at pos %d, used to appear at %d\n", i, lastAppear)
			startingLoc = max(startingLoc, lastAppear+1)
			fmt.Printf("Updating startloc to %d \n", startingLoc)
		}

		appearance[int(v)] = i
		currentRecord = i - startingLoc + 1
		maxRecord = max(maxRecord, currentRecord)
		fmt.Printf("Current record is %d\n", currentRecord)
	}
	fmt.Println("+++++")
	return maxRecord
}

func main() {
	fmt.Println(lengthOfLongestSubstring("defiance"))
}
