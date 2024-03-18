package medium

import "fmt"

// func pushToHeap(h *[]int, x int) {
// 	old := *h
// 	// 1,2,3,4,5 comes a 3
// 	insertLoc := len(old)
// 	// here we should do binary search
// 	for i := range old {
// 		if x < old[i] {
// 			insertLoc = i
// 			break
// 		}
// 	}
// 	fmt.Println("position is", insertLoc)
// 	new := []int{}
// 	new = append(new, old[0:insertLoc]...)
// 	new = append(new, x)
// 	new = append(new, old[insertLoc:]...)
// 	*h = new
// }

func pushToHeapBinarySearchPos(h *[]int, x int) {
	old := *h
	// 1,2,3,4,5 comes a 3

	// here we should do binary search

	n := len(old)

	left, right := 0, n
	for left < right {
		mid := (left + right) / 2
		midVal := old[mid]
		if x < midVal {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// fmt.Println("left is", left)
	// fmt.Println("right is", right)
	new := []int{}
	new = append(new, old[0:left]...)
	new = append(new, x)
	new = append(new, old[left:]...)
	*h = new
}

func popFromHeap(h *[]int) int {
	var x int
	old := *h
	// smallest at first
	x = old[0]
	// remove the first element
	*h = old[1:]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// lets have a not so optimized heap first
	// this heap is always sorted ascending
	heap := []int{}
	loc := 0
	n := len(heights)
	for {
		if loc+1 == n {
			return loc
		}

		diff := heights[loc+1] - heights[loc]

		if diff <= 0 {
			loc += 1
			continue
		}
		// we need to use ladder or bricks
		// now dif is > 0

		if ladders > 0 {
			ladders -= 1
			loc += 1
			pushToHeapBinarySearchPos(&heap, diff)
		} else {
			// we have used up the ladders

			if len(heap) > 0 && diff > heap[0] && bricks >= heap[0] {
				// least ladder usage too small, and we can use bricks for the old
				// and use ladder now
				leastLadderUsage := popFromHeap(&heap)

				bricks -= leastLadderUsage

				// ladders remain the same
				// then add diff to heap
				pushToHeapBinarySearchPos(&heap, diff)
				loc += 1
			} else {
				// simply use bricks
				if bricks >= diff {
					bricks -= diff
					loc += 1
				} else {
					return loc
				}
			}
		}

	}
	// return loc
}

func FurthestBuildingYouCanReach() {

	fmt.Println(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1))

}
