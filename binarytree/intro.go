package binarytree

import "fmt"

type justANode struct {
	val   string
	left  *justANode
	right *justANode
}

func prettyPrintBTree(r *justANode) {
	// lets try bfs
	// so we can print out line by line

	// queue is a slice of pointers!
	type justANodeWithLayer struct {
		node  *justANode
		layer int
	}
	queue := []*justANodeWithLayer{}
	src := &justANodeWithLayer{node: r}

	// maxLayer := 0
	// layerWiseMap := map[int][]string{}

	for {
		fmt.Println(src.node.val, "of layer", src.layer)
		leftNode, rightNode := src.node.left, src.node.right

		if leftNode == nil && rightNode == nil {
			// both nil.. do nothing
		} else if leftNode != nil && rightNode != nil {
			queue = append(queue, &justANodeWithLayer{node: leftNode, layer: src.layer + 1})
			queue = append(queue, &justANodeWithLayer{node: rightNode, layer: src.layer + 1})
		} else if leftNode == nil {
			queue = append(queue, &justANodeWithLayer{node: &justANode{val: " "}, layer: src.layer + 1})
			queue = append(queue, &justANodeWithLayer{node: rightNode, layer: src.layer + 1})
		} else if rightNode == nil {
			queue = append(queue, &justANodeWithLayer{node: leftNode, layer: src.layer + 1})
			queue = append(queue, &justANodeWithLayer{node: &justANode{val: " "}, layer: src.layer + 1})
		}

		if len(queue) == 0 {
			break
		}
		// bfs
		src, queue = queue[0], queue[1:]
		// dfs
		// r, queue = queue[len(queue)-1], queue[:len(queue)-1]
	}
}

func dfsRecursiveBinaryTree(r *justANode) {

	if r != nil {
		fmt.Println(r.val)
		dfsRecursiveBinaryTree(r.left)
		dfsRecursiveBinaryTree(r.right)
	}
}

func BinaryTreeIntro() {

	//          a
	//         / \
	//        b   c
	//       / \   \
	//      d   e   f
	//
	// a is the parent of b/c
	// b is the parent of d/e
	// c has only one child, f
	// a is the root, having no parent
	// leaf is d/e/f
	// for binary trees, every node has most 2 children!
	// one root only
	// exact one path from root to any node
	// no nodes is an empty tree, also a binary tree!

	fmt.Println("this is binary tree intro")

	a := justANode{val: "a"}
	b := justANode{val: "b"}
	c := justANode{val: "c"}
	d := justANode{val: "d"}
	e := justANode{val: "e"}
	f := justANode{val: "f"}

	fmt.Println(a, b, c, d, e, f)

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	prettyPrintBTree(&a)
	dfsRecursiveBinaryTree(&a)
}
