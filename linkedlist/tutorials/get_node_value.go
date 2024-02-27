package tutorials

import "fmt"

func getNodeValueRecursive(n *node, index int) any {
	if n == nil {
		return ""
	}
	if index == 0 {
		return n.val
	}
	// if node.val == target {
	// 	return true
	// }
	// return findFromLinkedList(node.next, target)
	return getNodeValueRecursive(n.next, index-1)
}

func getNodeValueIterative(n *node, index int) any {
	for n != nil {
		if index == 0 {
			return n.val
		}
		index -= 1
		n = n.next
	}
	return ""
}

func GetNodeValue() {
	strings := []string{"a", "v", "c", "f"}
	a := sliceToLinkedList(strings)
	fmt.Println(getNodeValueRecursive(a, 4))
	fmt.Println(getNodeValueIterative(a, 4))

	fmt.Println(getNodeValueRecursive(a, 3))
	fmt.Println(getNodeValueIterative(a, 3))
}
