package medium

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
}

func sliceToListNode(s []int) *listNode {
	startingNode := listNode{Val: s[0]}
	node := &startingNode
	for i, v := range s {
		if i > 0 {
			node.Next = &listNode{Val: v}
			node = node.Next
		}
	}
	return &startingNode
}

func printListNode(l *listNode) {
	n := l
	res := []int{}
	for {
		if n == nil {
			break
		}
		res = append(res, n.Val)
		n = n.Next
	}
	fmt.Println("list node:", res)
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {

	startingNode := listNode{}
	mainTravelNode := &startingNode
	l1TravelNode := l1
	l2TravelNode := l2

	var extra int = 0

	for {
		mainTravelNode.Val = l1TravelNode.Val + l2TravelNode.Val + extra
		if mainTravelNode.Val > 9 {
			extra = 1
			mainTravelNode.Val -= 10
		} else {
			extra = 0
		}

		if extra > 0 || l1TravelNode.Next != nil || l2TravelNode.Next != nil {
			mainTravelNode.Next = &listNode{}
			if l1TravelNode.Next == nil {
				l1TravelNode.Next = &listNode{}
			}
			if l2TravelNode.Next == nil {
				l2TravelNode.Next = &listNode{}
			}

			l1TravelNode = l1TravelNode.Next
			l2TravelNode = l2TravelNode.Next
			mainTravelNode = mainTravelNode.Next
		} else {
			break
		}
	}
	return &startingNode
}

func AddTwoNumbers() {
	// Input: l1 = [2,4,3], l2 = [5,6,4]
	// Output: [7,0,8]
	// Explanation: 342 + 465 = 807
	s1 := []int{9, 9, 9, 9, 9, 9}
	s2 := []int{9, 9, 9}

	l1 := sliceToListNode(s1)
	l2 := sliceToListNode(s2)

	printListNode(addTwoNumbers(l1, l2))
}
