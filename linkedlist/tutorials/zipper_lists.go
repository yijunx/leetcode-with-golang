package tutorials

import "fmt"

func zipperListsInterative(n1 *node, n2 *node) *node {
	resultNode := &node{}
	travellingNode := resultNode

	for n1 != nil && n2 != nil {
		nextOfN1 := n1.next
		nextOfN2 := n2.next

		// 0 -> a
		// 0 -> a -> *
		// 0 -> a -> * -> b
		// 0 -> a -> * -> b -> -

		travellingNode.next = n1
		travellingNode = travellingNode.next
		travellingNode.next = n2
		travellingNode = travellingNode.next
		// next round
		n1 = nextOfN1
		n2 = nextOfN2
	}

	// appending the remaining
	if n1 != nil {
		travellingNode.next = n1
	}
	// we dont need to consider n2, because
	// we do n2 second!! in the above for loop

	return resultNode.next
}

func zipperListsRecursive(n1 *node, n2 *node) *node {

	if n1 == nil && n2 == nil {
		// both exhausted
		return nil
	}

	if n1 == nil {
		// n1 exhausted
		return n2
	}

	if n2 == nil {
		// n2 exhausted
		return n1

	}
	head := n1

	// suppose we have
	// 1 2 3 4 5
	// a b c d e

	// 1 -> a -> next zipperListsRecursive{
	//     2 3 4 5
	//     b c d e
	// }

	// save the pointer
	nextN1 := n1.next
	nextN2 := n2.next
	// something like this
	n1.next = n2
	n2.next = zipperListsRecursive(nextN1, nextN2)

	return head
}

func ZipperLists() {
	string1 := []string{"a", "b", "c", "d", "e"}
	string2 := []string{"*", "-", "+"}
	a := sliceToLinkedList(string1)
	b := sliceToLinkedList(string2)
	fmt.Println(linkedListToSliceRecursive(a))
	fmt.Println(linkedListToSliceRecursive(b))

	fmt.Println(linkedListToSliceRecursive(zipperListsInterative(a, b)))

	a = sliceToLinkedList(string1)
	b = sliceToLinkedList(string2)
	fmt.Println(linkedListToSliceRecursive(zipperListsInterative(b, a)))

	a = sliceToLinkedList(string1)
	b = sliceToLinkedList(string2)
	fmt.Println(linkedListToSliceRecursive(zipperListsRecursive(a, b)))

	a = sliceToLinkedList(string1)
	b = sliceToLinkedList(string2)
	fmt.Println(linkedListToSliceRecursive(zipperListsRecursive(b, a)))

}
