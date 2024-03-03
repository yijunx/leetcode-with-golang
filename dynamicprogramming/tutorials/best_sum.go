package tutorials

import (
	"fmt"
)

func bestSum(target int, nums []int) []int {
	memoHow := map[int][]int{}
	// fmt.Println("using", nums, "to get", target)
	return bestSumRecursive(target, nums, memoHow)
}

func bestSumRecursive(target int, nums []int, memoHow map[int][]int) []int {
	if how, exists := memoHow[target]; exists {
		return how
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		// fmt.Println("well this target is hopeless")
		return nil
	}

	var leastWay []int = nil
	var currentBestLength int = target + 1

	for _, num := range nums {
		remainder := target - num
		if remainder < 0 {
			continue
		}
		// fmt.Println("num is", num, "remainder", remainder)
		howToGetRemainder := bestSumRecursive(remainder, nums, memoHow)

		if howToGetRemainder != nil {
			if len(howToGetRemainder) < currentBestLength {
				leastWay = []int{}
				// always safe to append to yourself!
				// avoid leastway = append(howToGetRemainder, num)
				leastWay = append(leastWay, num)
				leastWay = append(leastWay, howToGetRemainder...)
				currentBestLength = len(howToGetRemainder)
			}
		}

	}

	// fmt.Println(target, leastWay)
	memoHow[target] = leastWay
	return leastWay
}

func bestSumTabu(target int, nums []int) []int {
	dpHow := make([][]int, target+1)
	dpCan := make([]bool, target+1)

	dpCan[0] = true

	for currSum, currBestHow := range dpHow {
		if !dpCan[currSum] {
			continue
		}

		for _, num := range nums {
			if (currSum + num) <= target {

				// get curr how
				currHow := []int{}
				currHow = append(currHow, currBestHow...)
				currHow = append(currHow, num)
				// currHow

				// compare currHowWithBestHow
				if dpCan[currSum+num] {
					// been there before..
					// lets compare and replace if shorter
					if len(currHow) < len(dpHow[currSum+num]) {
						dpHow[currSum+num] = currHow
					}
				} else {
					dpHow[currSum+num] = currHow
				}
				// mark true
				dpCan[currSum+num] = true
			}
		}
	}
	return dpHow[target]
}

func BestSum() {
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(53, []int{50, 1}))
	fmt.Println(bestSum(299, []int{201, 200, 50, 1}))
	fmt.Println(bestSum(301, []int{201, 200, 50, 1}))
	fmt.Println(bestSum(301, []int{7, 14, 28}))

	fmt.Println(bestSumTabu(8, []int{2, 3, 5}))
	fmt.Println(bestSumTabu(53, []int{50, 1}))
	fmt.Println(bestSumTabu(299, []int{201, 200, 50, 1}))
	fmt.Println(bestSumTabu(301, []int{201, 200, 50, 1}))
	fmt.Println(bestSumTabu(301, []int{7, 14, 28}))
}
