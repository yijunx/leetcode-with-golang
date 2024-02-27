package tutorials

import "fmt"

func reverseLinkedListIterativeWithoutSlice(n *node) *node {
	resultNode := &node{}

	for n != nil {
		nextNode := n.next
		// target    node
		//      0 -> 1
		//      0 -> 2 -> 1
		tmp := resultNode.next
		resultNode.next = n
		n.next = tmp
		// can do it in one line
		// resultNode.next, node.next = node, resultNode.next

		n = nextNode
	}
	return resultNode.next
}

func reverseLinkedListIterative(n *node) *node {
	allNodes := []*node{}
	for n != nil {
		allNodes = append(allNodes, n)
		n = n.next
	}
	// reverse
	for i := len(allNodes) - 1; i >= 0; i-- {
		if i-1 >= 0 {
			allNodes[i].next = allNodes[i-1]
		} else {
			allNodes[i].next = nil
		}

	}
	return allNodes[len(allNodes)-1]
}

func reverseLinkedListRecursive(n *node) *node {
	resultNode := &node{}
	travelRecursive(n, resultNode)
	return resultNode.next
}

func travelRecursive(n *node, resultNode *node) {
	if n != nil {
		// here we first remember the next node
		next_node := n.next
		// 0
		// 0 - 1
		// 0 - 2 - 1
		// 0 - 3 - 2 - 1
		//tmp := resultNode.next
		fmt.Println("---iterations---")
		fmt.Println("making", resultNode, "->", n)
		fmt.Println("making", n, "->", resultNode.next)
		resultNode.next, n.next = n, resultNode.next
		//node.next = tmp
		travelRecursive(next_node, resultNode)
	}
}

// a b c d
//

func ReverseLinkedList() {
	numbers := []int{1, 2, 3, 4, 5}
	a := sliceToLinkedList(numbers)
	fmt.Println(linkedListToSliceRecursive(reverseLinkedListIterative(a)))
	// we have to regenerated it, because
	a = sliceToLinkedList(numbers)
	fmt.Println(linkedListToSliceRecursive(reverseLinkedListRecursive(a)))
	a = sliceToLinkedList(numbers)
	fmt.Println(linkedListToSliceRecursive(reverseLinkedListIterativeWithoutSlice(a)))
}
