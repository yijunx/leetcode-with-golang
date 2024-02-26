package tutorials

import "fmt"

func getNodeValueRecursive(node *justANode, index int) string {
	if node == nil {
		return ""
	}
	if index == 0 {
		return node.val
	}
	// if node.val == target {
	// 	return true
	// }
	// return findFromLinkedList(node.next, target)
	return getNodeValueRecursive(node.next, index-1)
}

func getNodeValueIterative(node *justANode, index int) string {
	for node != nil {
		if index == 0 {
			return node.val
		}
		index -= 1
		node = node.next
	}
	return ""
}

func GetNodeValue() {
	strings := []string{"a", "v", "c", "f"}
	a := stringSliceToLinkedList(strings)
	fmt.Println(getNodeValueRecursive(a, 4))
	fmt.Println(getNodeValueIterative(a, 4))

	fmt.Println(getNodeValueRecursive(a, 3))
	fmt.Println(getNodeValueIterative(a, 3))
}
