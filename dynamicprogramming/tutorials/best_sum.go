package tutorials

import (
	"fmt"
)

func bestSum(target int, nums []int) []int {
	memoHow := map[int][]int{}
	fmt.Println("using", nums, "to get", target)
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
func BestSum() {
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	// fmt.Println(bestSum(9, []int{50, 1}))
	fmt.Println(bestSum(53, []int{50, 1}))
	// fmt.Println(bestSum(10, []int{2, 3, 4, 1}))
	// fmt.Println(bestSum(23, []int{5, 8}))
	// fmt.Println(bestSum(51, []int{2, 3, 4, 17, 48}))
	// fmt.Println(bestSum(299, []int{201, 200, 50, 1}))
	// fmt.Println(bestSum(300, []int{200, 10, 1}))
	// fmt.Println(bestSum(300, []int{7, 14, 28}))
	// fmt.Println(bestSum(280, []int{279, 7, 14, 56, 1}))

	// s1 := make([]int, 3, 4)
	// s2 := s1[:2]
	// fmt.Println(s1, s2)

	// s2[0]++
	// fmt.Println(s1, s2)
	// fmt.Println(s1[0] == s2[0]) //true

	// s1 = append(s1, 0)
	// s2[0]++
	// fmt.Println(s1, s2)
	// fmt.Println(s1[0] == s2[0]) //true

	// s1 = append(s1, 0)
	// s2[0]++
	// fmt.Println(s1, s2)
	// fmt.Println(s1[0] == s2[0]) //false
}
