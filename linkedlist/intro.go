package linkedlist

import "fmt"

type justANode struct {
	val  string
	next *justANode
}

func linkedListTraversalIterative(node *justANode) {
	for node != nil {
		fmt.Println(node.val)
		node = node.next
	}
}

func linkedListTraversalRecursive(node *justANode) {
	if node != nil {
		fmt.Println(node.val)
		linkedListTraversalRecursive(node.next)
	}
}

func LinkedListIntro() {
	// a -> b -> c -> d
	// head           tail
	// this is an ordered data structure

	a := justANode{val: "a"}
	b := justANode{val: "b"}
	c := justANode{val: "c"}
	d := justANode{val: "d"}

	a.next = &b
	b.next = &c
	c.next = &d

	linkedListTraversalIterative(&a)
	linkedListTraversalRecursive(&a)

	// what is the diffeence of array vs linked list

	// a -> b -> c -> d
	// head           tail
	// this is an ordered data structure

	// array, continuous in memory!
	// 0 1 2 3
	// a b c d

	// lets say of inserting at index 2 in an array
	// needs to move many stuff in an array!
	// each shift is one operation, thus such inserting (in the middle) has
	// time complexity of n

	// but linked insersion is constant time complexity

}
