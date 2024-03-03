package tutorials

import "fmt"

func canSum(target int, nums []int) bool {
	// well we can do it in a dfs way
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	// result := false

	// here we start the dfs
	for _, num := range nums {
		if canSum(target-num, nums) {
			return true
		}
		// result = result || canSum(target-num, nums)
	}
	return false
}

func canSum2(target int, nums []int) bool {
	// well we can do it in a dfs way
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	result := false

	// here we start the dfs
	for _, num := range nums {
		result = result || canSum(target-num, nums)
	}
	return result
}

func canSumWithMemo(target int, nums []int) bool {
	memo := map[int]bool{0: true}
	return canSumWithMemoDFS(target, nums, memo)
}

func canSumWithMemoDFS(target int, nums []int, memo map[int]bool) bool {
	if target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	// if result, found := memo[target]; found {
	// 	return result
	// }
	// well there is duplication of writing into memory!!

	// better to check / input the memory inside the dfs!!!
	for _, num := range nums {
		remainder := target - num

		if result, found := memo[remainder]; found {
			return result
		}

		new_result := canSumWithMemoDFS(remainder, nums, memo)
		// save the result
		fmt.Println("saving", remainder, "to memo as", new_result)
		memo[remainder] = new_result
		if new_result {
			return true
		}
	}

	return false
}

func canSumTabu(target int, nums []int) bool {
	// init the dp
	dp := make([]bool, target+1)
	// base case
	dp[0] = true

	// now we have a dp,
	for i, val := range dp {
		if !val {
			continue
		}
		for _, num := range nums {
			if i+num <= target {
				dp[i+num] = true
			}
		}
	}
	return dp[target]
}

func CanSum() {
	fmt.Println(canSum(23, []int{5, 8, 1}))
	fmt.Println(canSum2(23, []int{5, 8, 1}))
	fmt.Println(canSumWithMemo(23, []int{5, 8, 1}))
	fmt.Println(canSumTabu(23, []int{5, 8, 1}))

	fmt.Println(canSum(23, []int{2, 4, 6, 8, 10}))
	fmt.Println(canSum2(23, []int{2, 4, 6, 8, 10}))
	fmt.Println(canSumWithMemo(23, []int{2, 4, 6, 8, 10}))
	fmt.Println(canSumTabu(23, []int{2, 4, 6, 8, 10}))

	// fmt.Println(canSum(300, []int{7, 14})) // never ends..
	// fmt.Println(canSum2(300, []int{7, 14})) // never ends
	fmt.Println(canSumWithMemo(300, []int{7, 14})) // never ends..
	fmt.Println(canSumTabu(300, []int{7, 14}))     // never ends..
}
