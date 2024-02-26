package tutorials

import "fmt"

func sumOfLinkedListRecursive(node *numberNode) int {
	if node == nil {
		return 0
	} else {
		return sumOfLinkedListRecursive(node.next) + node.val
	}
}

func SumOfLinkedList() {
	numbers := []int{1, 2, 3, 4}
	a := sliceToLinkedList(numbers)

	fmt.Println(sumOfLinkedListRecursive(a))
}
