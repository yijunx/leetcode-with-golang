package tutorials

import "fmt"

func howSum(target int, nums []int) []int {
	memoHow := map[int][]int{}
	memoCan := map[int]bool{}
	how, _ := howSumRecursive(target, nums, memoHow, memoCan)
	return how
}

func howSumRecursive(target int, nums []int, memoHow map[int][]int, memoCan map[int]bool) ([]int, bool) {
	if target == 0 {
		// which is good, returns an empty list,
		// so later we can append
		return []int{}, true
	}
	if target < 0 {
		return []int{}, false
	}

	for _, num := range nums {

		remainder := target - num

		// check if we knnow how to solve the reminder
		if can, exists := memoCan[remainder]; exists {
			if can {
				memoHow[target] = append(memoHow[remainder], num)
				return memoHow[target], true
			} else {
				return []int{}, false
			}
		}

		new_how, new_can := howSumRecursive(remainder, nums, memoHow, memoCan)
		// fmt.Println("saving", remainder, "to memo as", new_can)
		memoCan[remainder] = new_can
		// check if we know how to deal with the reminder
		if new_can {
			// memoHow[remainder] = new_how
			memoHow[target] = append(new_how, num)
			return memoHow[target], true
		}

	}
	return []int{}, false
}

func howSumTabu(target int, nums []int) []int {
	// declare
	dpHow := make([][]int, target+1)
	dpCan := make([]bool, target+1)

	dpCan[0] = true

	// initialize
	// for i := range dp {
	// 	dp[i] = []int{}
	// }
	// dp[0] = []int{}

	// for _, fff := range dp {
	// 	fff = []int{}
	// }

	// loop
	for currentTarget, howToAdd := range dpHow {
		if !dpCan[currentTarget] {
			continue
		}
		for _, num := range nums {
			if currentTarget+num <= target {
				if !dpCan[currentTarget+num] {
					// if we have not arrived here before!!!
					dpCan[currentTarget+num] = true
					dpHow[currentTarget+num] = append(dpHow[currentTarget+num], howToAdd...)
					dpHow[currentTarget+num] = append(dpHow[currentTarget+num], num)
				}
			}
		}
	}
	return dpHow[target]
}

func HowSum() {
	fmt.Println(howSum(23, []int{5, 8}))
	fmt.Println(howSum(51, []int{2, 4}))
	fmt.Println(howSum(300, []int{7, 14}))
	fmt.Println(howSum(280, []int{7, 14}))
	fmt.Println(howSum(5, []int{1,2,5}))

	fmt.Println(howSumTabu(23, []int{5, 8}))
	fmt.Println(howSumTabu(51, []int{2, 4}))
	fmt.Println(howSumTabu(300, []int{7, 14}))
	fmt.Println(howSumTabu(280, []int{7, 14}))
}
