package tutorials

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func printATree(root *treeNode) {

	// lets use bfs
	// for bfs, we need to queue
	queue := []*treeNode{}
	src := root

	for {
		if src != nil {
			fmt.Println(src.val)
			queue = append(queue, src.left)
			queue = append(queue, src.right)
		}

		if len(queue) == 0 {
			break
		}
		src, queue = queue[0], queue[1:]
	}

}

func insertValIntoBST(root *treeNode, value int) {
	if value > root.val {
		if root.right == nil {
			root.right = &treeNode{val: value}
		} else {
			insertValIntoBST(root.right, value)
		}
	} else {
		if root.left == nil {
			root.left = &treeNode{val: value}
		} else {
			insertValIntoBST(root.left, value)
		}
	}
}

func insertValIntoBST2(root *treeNode, value int) {
	if root == nil {
		root = &treeNode{val: value}
	}
	if value > root.val {
		insertValIntoBST(root.right, value)
	} else {
		insertValIntoBST(root.left, value)
	}
}

func InsertingValueIntoBST() {

	//          50
	//         /  \
	//        25   75
	//       /  \    \
	//      10   30   80

	a := treeNode{val: 50}
	b := treeNode{val: 25}
	c := treeNode{val: 75}
	d := treeNode{val: 10}
	e := treeNode{val: 30}
	f := treeNode{val: 80}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	printATree(&a)

	insertValIntoBST(&a, 1)
	insertValIntoBST2(&a, 2)

	fmt.Println("after inserting")

	printATree(&a)

}
