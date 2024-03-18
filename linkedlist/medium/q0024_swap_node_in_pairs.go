package medium

import "fmt"

func swapPairs(head *listNode) *listNode {

	resultHead := &listNode{}
	resultTrav := resultHead
	resultReve := resultHead

	appendReverse := false
	// comes
	// 1 2 3 4 5
	// 0
	// resultHead    resultTrav   head becomes
	// 0 -> 1                     2 -> 3 -> 4 -> 5 ... to append 1: result rev at 0, result trav at 0
	// 0 -> 2 -> 1                3 -> 4 -> 5 ...      to append 2: result rev at 0, result trav at 1
	// 0 -> 2 -> 1 -> 3           4 -> 5
	// 0 -> 2 -> 1 -> 4 -> 3      5
	// 0 -> 2 -> 1 -> 4 -> 3 -> 5 nil

	for head != nil {
		// save the head
		fmt.Println("##### we see", head.Val)
		theNext := head.Next

		if appendReverse {
			// fmt.Println("append reversely", head.Val)
			resultReve.Next, head.Next = head, resultReve.Next
			resultReve = resultReve.Next.Next
			resultTrav = resultReve
			printListNode(resultHead)
			fmt.Println("resultReve is now at", resultReve.Val)
			fmt.Println("resultTrav is now at", resultTrav.Val)
		} else {
			// fmt.Println("append normally", head.Val)
			resultTrav.Next, head.Next = head, resultTrav.Next
			resultTrav = resultTrav.Next
			printListNode(resultHead)
			fmt.Println("resultReve is now at", resultReve.Val)
			fmt.Println("resultTrav is now at", resultTrav.Val)
		}
		head = theNext
		appendReverse = !appendReverse
	}
	return resultHead.Next
}

func SwapPairs() {

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := []int{1}

	l1 := sliceToListNode(s1)
	l2 := sliceToListNode(s2)

	h1 := swapPairs(l1)
	h2 := swapPairs(l2)
	//
	printListNode(h1)
	printListNode(h2)
	// printListNode(swapPairs(l1))
	// printListNode(swapPairs(l2))
}
