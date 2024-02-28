package tutorials

import "fmt"

type listNode struct {
	val  any
	next *listNode
}

func sliceToLinkedList[T any](list []T) *listNode {
	travelingNode := &listNode{}
	head := travelingNode
	for _, val := range list {
		travelingNode.next = &listNode{val: val}
		travelingNode = travelingNode.next
	}
	return head.next
}

func linkedListToSlice(node *listNode) []any {
	result := []any{}

	for node != nil {
		result = append(result, node.val)
		node = node.next
	}
	return result
}

func reverseLinkedList(node *listNode) *listNode {
	resultNode := &listNode{}
	reverseLinkedListRecursive(node, resultNode)
	return resultNode.next
}

func reverseLinkedListRecursive(node *listNode, resultNode *listNode) {
	// 0 1
	// 0 2 1
	// 0 3 2 1
	if node != nil {
		nextNode := node.next
		resultNode.next, node.next = node, resultNode.next
		reverseLinkedListRecursive(nextNode, resultNode)
	}
}

func ReverseLinkedList() {
	a := []int{1, 2, 3, 4, 5}
	node := sliceToLinkedList(a)
	fmt.Println(linkedListToSlice(node))
	fmt.Println(linkedListToSlice(reverseLinkedList(node)))
}
