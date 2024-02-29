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
			// how_to_add_to_target := append(how_to_add_reminder, num)
			// fmt.Println("saving", how_to_add_reminder, "to target:", remainder)
			memoHow[remainder] = new_how
			memoHow[target] = append(new_how, num)
			return memoHow[target], true
		}

	}
	return []int{}, false
}
func HowSum() {
	fmt.Println(howSum(23, []int{5, 8}))
	fmt.Println(howSum(51, []int{2, 4}))
	fmt.Println(howSum(300, []int{7, 14}))
}
