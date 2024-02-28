package tutorials

import "fmt"

func binarySearch(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return target == nums[0]
	}
	midPoint := len(nums) / 2

	return binarySearch(nums[0:midPoint], target) || binarySearch(nums[midPoint:], target)
}

func binarySearchFindIndex(nums []int, target int) int {
	left, right := 0, len(nums)
	return binarySearchWithLeftAndRight(nums, target, left, right)
}

func binarySearchWithLeftAndRight(nums []int, target int, left int, right int) int {

	if left > right {
		return -1
	}
	mid := (left + right) / 2

	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		right = mid - 1
	}
	if nums[mid] < target {
		left = mid + 1
	}

	return binarySearchWithLeftAndRight(nums, target, left, right)
}

func BinarySearch() {
	nums := []int{1, 4, 5, 10, 29, 100, 234, 456, 12345, 99023, 123456}

	fmt.Println(binarySearch(nums, 5))
	fmt.Println(binarySearch(nums, 123456))
	fmt.Println(binarySearch(nums, 11))

	fmt.Println(binarySearchFindIndex(nums, 5))
	fmt.Println(binarySearchFindIndex(nums, 123456))
	fmt.Println(binarySearchFindIndex(nums, 11))
}
