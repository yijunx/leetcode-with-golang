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
	resultNode := &node{}
	travellingNode := resultNode
	exploreLinkedListWithTravelingNode(n1, n2, travellingNode)
	return resultNode.next
}

func exploreLinkedListWithTravelingNode(n1 *node, n2 *node, t *node) {
	if n1 != nil && n2 != nil {
		tempN1 := n1.next
		tempN2 := n2.next
		t.next = n1
		t = t.next
		t.next = n2
		t = t.next
		exploreLinkedListWithTravelingNode(tempN1, tempN2, t)
	}
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
