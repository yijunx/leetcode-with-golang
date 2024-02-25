package tutorials

import "fmt"

func treeMaxPathSumBFS(node *numberNode) int {

	type numberNodeWithTotal struct {
		node  *numberNode
		total int
	}
	nodeWithTotal := &numberNodeWithTotal{
		node:  node,
		total: 0, // the total before reaching here
	}
	currentMax := 0
	queue := []*numberNodeWithTotal{}
	for {
		if nodeWithTotal.node != nil {
			currentTillHere := nodeWithTotal.node.val + nodeWithTotal.total
			currentMax = max(currentMax, currentTillHere)
			queue = append(queue, &numberNodeWithTotal{node: nodeWithTotal.node.left, total: currentTillHere})
			queue = append(queue, &numberNodeWithTotal{node: nodeWithTotal.node.right, total: currentTillHere})
		}

		if len(queue) == 0 {
			break
		}

		nodeWithTotal, queue = queue[0], queue[1:]
	}
	return currentMax
}

func treeMaxPathSumDFS(node *numberNode) int {
	if node != nil {
		return node.val + max(treeMaxPathSumDFS(node.left), treeMaxPathSumDFS(node.right))
	}
	return 0
}

func TreeMaxPathSum() {

	a := numberNode{val: 100}
	b := numberNode{val: 23}
	c := numberNode{val: 3}
	d := numberNode{val: 4}
	e := numberNode{val: 5}
	f := numberNode{val: 6}

	//          a
	//         / \
	//        b   c
	//       / \   \
	//      d   e   f

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	fmt.Println(treeMaxPathSumBFS(&a))
	fmt.Println(treeMaxPathSumDFS(&a))
}
