package tutorials

import "fmt"

func sliceToLinkedList[T any](anies []T) *node {
	firstVal, anies := anies[0], anies[1:]
	n := &node{
		val: firstVal,
	}
	headNode := n
	for _, val := range anies {
		newNode := &node{
			val: val,
		}
		n.next = newNode
		n = newNode
	}
	return headNode
}

// func stringSliceToLinkedList(strings []string) *stringNode {
// 	firstVal, strings := strings[0], strings[1:]
// 	node := &stringNode{
// 		val: firstVal,
// 	}
// 	headNode := node
// 	for _, val := range strings {
// 		newNode := &stringNode{
// 			val: val,
// 		}
// 		node.next = newNode
// 		node = newNode
// 	}
// 	return headNode
// }

func linkedListToSliceIterative(n *node) []any {
	vals := []any{}
	for n != nil {
		vals = append(vals, n.val)
		n = n.next
	}
	return vals
}

func linkedListToSliceRecursive(n *node) []any {
	result := []any{}
	exploreLinkedListWithResult(n, &result)
	return result
}

func exploreLinkedListWithResult(n *node, result *[]any) {
	// just append things into result, return nothing
	if n == nil {
		return
	}
	*result = append(*result, n.val)
	exploreLinkedListWithResult(n.next, result)
}

func LinkedListToSlice() {
	numbers := []int{1, 2, 3, 4}
	a := sliceToLinkedList(numbers)

	fmt.Println(linkedListToSliceIterative(a))
	fmt.Println(linkedListToSliceRecursive(a))
}
