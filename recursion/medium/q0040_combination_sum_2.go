package medium

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	fmt.Println(candidates)
	solve(candidates, target, 0, []int{}, &result)
	return result
}

func solve(nums []int, target int, index int, lst []int, result *[][]int) {
	fmt.Println("tartget is", target, "index is", index)
	fmt.Println("lst are", lst)
	if target == 0 {
		fmt.Println("####### result appending", lst)
		newlst := []int{}
		newlst = append(newlst, lst...)
		*result = append(*result, newlst)
		return
	}

	if index == len(nums) || target < nums[index] {
		return
	}

	candidate := nums[index]
	lst = append(lst, candidate)
	// Recur with reduced target and move to next index
	solve(nums, target-candidate, index+1, lst, result)
	// pop
	lst = lst[:len(lst)-1]

	i := 1
	for index+i < len(nums) && nums[index+i] == candidate {
		i = i + 1
	} // skipping same numbers

	// Recur skip the current!
	fmt.Println("before check: lst is", lst, "index+i is", index+i, "target is", target)
	solve(nums, target, index+i, lst, result)
}

func CombinationSum2() {
	// the key is we cannot have replicates in the answer..
	candidates := []int{3, 1, 3, 5, 1, 1}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
