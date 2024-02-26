package tutorials

import "fmt"

func sliceToLinkedList(numbers []int) *numberNode {
	firstNumber, numbers := numbers[0], numbers[1:]
	node := &numberNode{
		val: firstNumber,
	}
	headNode := node
	for _, number := range numbers {
		newNode := &numberNode{
			val: number,
		}
		node.next = newNode
		node = newNode
	}
	return headNode
}

func stringSliceToLinkedList(strings []string) *justANode {
	firstVal, strings := strings[0], strings[1:]
	node := &justANode{
		val: firstVal,
	}
	headNode := node
	for _, val := range strings {
		newNode := &justANode{
			val: val,
		}
		node.next = newNode
		node = newNode
	}
	return headNode
}

func linkedListToSliceIterative(node *numberNode) []int {
	numbers := []int{}
	for node != nil {
		numbers = append(numbers, node.val)
		node = node.next
	}
	return numbers
}

func linkedListToSliceRecursive(node *numberNode, numbers []int) []int {
	if node != nil {
		// tail first
		// return append(linkedListToSliceRecursive(node.next, numbers), node.val)
		// head first, seems not so efficent??
		return append([]int{node.val}, linkedListToSliceRecursive(node.next, numbers)...)
	} else {
		return numbers
	}
}

func LinkedListToSliceRecursive2(node *numberNode) []int {
	result := []int{}
	exploreLinkedListWithResult(node, &result)
	return result
}

func exploreLinkedListWithResult(node *numberNode, result *[]int) {
	// just append things into result, return nothing
	if node == nil {
		return
	}
	*result = append(*result, node.val)
	exploreLinkedListWithResult(node.next, result)
}

func LinkedListToSlice() {
	numbers := []int{1, 2, 3, 4}
	a := sliceToLinkedList(numbers)

	fmt.Println(linkedListToSliceIterative(a))
	fmt.Println(linkedListToSliceRecursive(a, []int{}))
	fmt.Println(LinkedListToSliceRecursive2(a))
}
