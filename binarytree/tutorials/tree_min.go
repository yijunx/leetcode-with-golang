package tutorials

import "fmt"

func treeMinBFS(node *numberNode) int {
	const MaxUint = ^uint(0)
	// const MinUint = 0
	var currMin = int(MaxUint >> 1)
	// const MinInt = -MaxInt - 1

	queue := []*numberNode{}
	for {
		if node != nil {
			currMin = min(currMin, node.val)
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}

		if len(queue) == 0 {
			break
		}

		node, queue = queue[0], queue[1:]
	}

	return currMin
}

func treeMinDFS(node *numberNode) int {
	if node == nil {
		return 65535
	} else {
		return min(node.val, treeMinDFS(node.left), treeMinDFS(node.right))
	}
}

func TreeMin() {

	a := numberNode{val: 100}
	b := numberNode{val: 23}
	c := numberNode{val: 3}
	d := numberNode{val: 4}
	e := numberNode{val: 5}
	f := numberNode{val: 6}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	fmt.Println(treeMinBFS(&a))
	fmt.Println(treeMinDFS(&a))
}
