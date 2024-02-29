package tutorials

import "fmt"

type numberNode struct {
	val  int
	next *numberNode
}

func listToNumberNodes(list []int) *numberNode {
	result := &numberNode{}
	travelingNode := result
	for _, v := range list {
		travelingNode.next = &numberNode{val: v}
		travelingNode = travelingNode.next
	}
	return result.next
}

func printNumberNodes(n *numberNode) {
	result := []int{}
	for n != nil {
		result = append(result, n.val)
		n = n.next
	}
	fmt.Println(result)
}

func mergeSortedLinkedList(node1, node2 *numberNode) *numberNode {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node1.val < node2.val {
		node1.next = mergeSortedLinkedList(node1.next, node2)
		return node1
	} else {
		node2.next = mergeSortedLinkedList(node1, node2.next)
		return node2
	}
}

func MergeSortedLinkedList() {
	nums1 := []int{1, 3, 5, 7, 9, 109}
	nums2 := []int{1, 5, 10, 14, 22, 104}

	node1 := listToNumberNodes(nums1)
	node2 := listToNumberNodes(nums2)

	printNumberNodes(mergeSortedLinkedList(node1, node2))
}
