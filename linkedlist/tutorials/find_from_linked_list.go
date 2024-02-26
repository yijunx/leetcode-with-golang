package tutorials

import "fmt"

func findFromLinkedList(node *justANode, target string) bool {
	if node == nil {
		return false
	}
	if node.val == target {
		return true
	}
	return findFromLinkedList(node.next, target)
}

func FindFromLinkedList() {
	numbers := []string{"a", "v", "c", "f"}
	a := stringSliceToLinkedList(numbers)

	fmt.Println(findFromLinkedList(a, "b"))
	fmt.Println(findFromLinkedList(a, "f"))
}
