package tutorials

import "fmt"

func findFromLinkedList(n *node, target string) bool {
	if n == nil {
		return false
	}
	if n.val == target {
		return true
	}
	return findFromLinkedList(n.next, target)
}

func FindFromLinkedList() {
	numbers := []string{"a", "v", "c", "f"}
	a := sliceToLinkedList(numbers)

	fmt.Println(findFromLinkedList(a, "b"))
	fmt.Println(findFromLinkedList(a, "f"))
}
