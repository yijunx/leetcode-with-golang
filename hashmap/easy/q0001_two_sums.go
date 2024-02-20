package easy

import "fmt"

func twoSum(nums []int, target int) []int {
	var intMap map[int]int = make(map[int]int)
	for i, v := range nums {
		var differenceIndex, found = intMap[v]

		if found {
			return []int{differenceIndex, i}
		}
		intMap[target-v] = i
	}
	return []int{}
}

func TwoSum() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}
