package tutorials

import "fmt"

func mergeSort(nums []int) []int {
	// we need to keep dividing
	// sort the smallest element
	// them put them back together
	mid := len(nums) / 2
	nums1 := nums[:mid]
	nums2 := nums[mid:]
	return partlyMergeSort(nums1, nums2)
}

func mergeSortedList(nums1 []int, nums2 []int) []int {
	index1 := 0
	index2 := 0
	result := []int{}

	for index1 < len(nums1) && index2 < len(nums2) {
		if nums1[index1] < nums2[index2] {
			result = append(result, nums1[index1])
			index1 += 1
		} else {
			result = append(result, nums2[index2])
			index2 += 1
		}
	}

	// appending the rest
	if index1 < len(nums1) {
		result = append(result, nums1[index1:]...)
	}
	if index2 < len(nums2) {
		result = append(result, nums2[index2:]...)
	}
	return result
}

func partlyMergeSort(nums1 []int, nums2 []int) []int {

	// sort most basic element
	lenNum1 := len(nums1)
	lenNum2 := len(nums2)

	// merge, put nums1 and num2 together
	if lenNum1 <= 2 && lenNum2 <= 2 {
		if lenNum1 == 2 {
			if nums1[1] < nums1[0] {
				nums1[1], nums1[0] = nums1[0], nums1[1]
			}
		}

		if lenNum2 == 2 {
			if nums2[1] < nums2[0] {
				nums2[1], nums2[0] = nums2[0], nums2[1]
			}
		}
		return mergeSortedList(nums1, nums2)
	} else {
		// keep dividing
		mid1 := lenNum1 / 2
		numsleft1 := nums1[:mid1]
		numsleft2 := nums1[mid1:]
		mid2 := lenNum2 / 2
		numsright1 := nums2[:mid2]
		numsright2 := nums2[mid2:]
		return mergeSortedList(partlyMergeSort(numsleft1, numsleft2), partlyMergeSort(numsright1, numsright2))
	}

}

func mergeSortAnotherWay(nums []int) []int {
	start, end := 0, len(nums)
	partlyMergeSortAnother(nums, start, end)
	return nums
}

func partlyMergeSortAnother(nums []int, start int, end int) {
	if start < end-1 {
		mid := (start + end) / 2
		partlyMergeSortAnother(nums, start, mid)
		partlyMergeSortAnother(nums, mid, end)
		merge(nums, start, end, mid)
	}
}

func merge(nums []int, start int, end int, mid int) {
	result := []int{}
	i := start
	j := mid

	for i < mid && j < end {
		if nums[i] < nums[j] {
			result = append(result, nums[i])
			i++
		} else {
			result = append(result, nums[j])
			j++
		}
	}

	if i < mid {
		result = append(result, nums[i:mid]...)
	}
	if j < end {
		result = append(result, nums[j:end]...)
	}

	// copy back
	// for index, value := range result {
	// 	nums[index] = value
	// }

	// well umm this is the key
	// well rememner result does not start from zero position of the nums!!!!
	for i = start; i < end; i++ {
		nums[i] = result[i-start]
	}
}

func playWithBinarySearch(nums []int, start int, end int) {
	fmt.Println("start is", start, "end is ", end)
	fmt.Println(nums[start:end])
	if start < end-1 {
		mid := (start + end) / 2
		fmt.Println("mid is", mid)
		playWithBinarySearch(nums, start, mid)
		playWithBinarySearch(nums, mid, end)
	}
}

func MergeSort() {
	a := []int{4, 1, 3, 2, 0, -1, 7, 10, 9, 20, 0, 3, 4, 1}
	fmt.Println(mergeSort(a))

	a = []int{4, 1, 3, 2, 0, -1, 7, 10, 9, 20, 0, 3, 4, 1}

	// x := []int{1, 4, 6, 10, 2, 7, 9, 20}
	// fmt.Println(len(x))
	// fmt.Println(x[0:4])
	// fmt.Println(x[4:8])
	// merge(x, 0, 7, 4)
	// fmt.Println(x)

	y := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	playWithBinarySearch(y, 0, len(y))
	fmt.Println(mergeSortAnotherWay(a))

	// fmt.Println(mergeSortedList([]int{1, 3, 5}, []int{2, 4, 6, 8}))

}
