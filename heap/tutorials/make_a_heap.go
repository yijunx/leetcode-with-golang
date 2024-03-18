package tutorials

// not done yet

import "fmt"

type heapNode struct {
	// we need to make sure that only the smallest val is on the top
	val   int
	left  *heapNode
	right *heapNode
}

func addToHeapNode(h *heapNode, x int) *heapNode {
	if h == nil {
		return &heapNode{val: x}
	}

	if x <= h.val {
		newh := &heapNode{val: x}
		newh.left = h
		h = newh
	} else {
		// bigger than val
		addToHeapNode(h.right, x)
	}
	return h
}

// func popFromHeapNode(h *heapNode) int {
// 	return 0
// }

func peakHeapNode(h *heapNode) int {
	return h.val
}

func MakeAHeap() {
	var h *heapNode

	addToHeapNode(h, 6)
	fmt.Println(peakHeapNode(h))
	// addToHeapNode(h, 3)
	// fmt.Println(peakHeapNode(h))
	// addToHeapNode(h, 4)
	// fmt.Println(peakHeapNode(h))
	// addToHeapNode(h, 1)
	// fmt.Println(peakHeapNode(h))
	// addToHeapNode(h, 9)
	// fmt.Println(peakHeapNode(h))
	// addToHeapNode(h, 5)
	// fmt.Println(peakHeapNode(h))
}

//
//
//
//
