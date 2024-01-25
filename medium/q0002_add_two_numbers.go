package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(s []int) *ListNode {
	startingNode := ListNode{Val: s[0]}
	node := &startingNode
	for i, v := range s {
		if i > 0 {
			node.Next = &ListNode{Val: v}
			node = node.Next
		}
	}
	return &startingNode
}

func PrintListNode(l *ListNode) {
	fmt.Println("Printing a list node")
	for {
		if l == nil {
			break
		}
		fmt.Println(l.Val)
		l = l.Next
	}
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	startingNode := ListNode{}
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
			mainTravelNode.Next = &ListNode{}
			if l1TravelNode.Next == nil {
				l1TravelNode.Next = &ListNode{}
			}
			if l2TravelNode.Next == nil {
				l2TravelNode.Next = &ListNode{}
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

// func main() {
// 	// Input: l1 = [2,4,3], l2 = [5,6,4]
// 	// Output: [7,0,8]
// 	// Explanation: 342 + 465 = 807
// 	s1 := []int{9, 9, 9, 9, 9, 9}
// 	s2 := []int{9, 9, 9}

// 	l1 := SliceToListNode(s1)
// 	l2 := SliceToListNode(s2)

// 	PrintListNode(l1)

// 	PrintListNode(AddTwoNumbers(l1, l2))
// }
