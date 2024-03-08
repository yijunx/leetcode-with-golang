package tutorials

import "fmt"

type treeNode struct {
	val   string
	left  *treeNode
	right *treeNode
}

func prettyPrintBTree(r *treeNode) {
	// lets try bfs
	// so we can print out line by line

	// queue is a slice of pointers!
	type justANodeWithLayer struct {
		node  *treeNode
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
			queue = append(queue, &justANodeWithLayer{node: &treeNode{val: " "}, layer: src.layer + 1})
			queue = append(queue, &justANodeWithLayer{node: rightNode, layer: src.layer + 1})
		} else if rightNode == nil {
			queue = append(queue, &justANodeWithLayer{node: leftNode, layer: src.layer + 1})
			queue = append(queue, &justANodeWithLayer{node: &treeNode{val: " "}, layer: src.layer + 1})
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

func dfsRecursiveBinaryTree(r *treeNode) {

	if r != nil {
		fmt.Println(r.val)
		dfsRecursiveBinaryTree(r.left)
		dfsRecursiveBinaryTree(r.right)
	}
}

func preorderPrint(r *treeNode) {
	if r != nil {
		fmt.Println(r.val)
		preorderPrint(r.left)
		preorderPrint(r.right)
	}
}

func preorderToSlice(r *treeNode, list *[]string) {
	if r != nil {
		*list = append(*list, r.val)
		preorderToSlice(r.left, list)
		preorderToSlice(r.right, list)
	}
}

func postorderPrint(r *treeNode) {
	if r != nil {
		postorderPrint(r.left)
		postorderPrint(r.right)
		fmt.Println(r.val)
	}
}

func postorderToSlice(r *treeNode, list *[]string) {
	if r != nil {
		postorderToSlice(r.left, list)
		postorderToSlice(r.right, list)
		*list = append(*list, r.val)
	}
}

func inorderPrint(r *treeNode) {
	if r != nil {
		inorderPrint(r.left)
		fmt.Println(r.val)
		inorderPrint(r.right)
	}
}

func inorderPrintWithoutRecurr(r *treeNode) {

	queue := []*treeNode{}
	for {

		// add stuff to queue
		if len(queue) == 0 {
			break
		}

		// here we need to decide if we can pop
		r = queue[0]
		// shorten the queue

	}
}

func inorderToSlice(r *treeNode, list *[]string) {
	if r != nil {
		inorderToSlice(r.left, list)
		*list = append(*list, r.val)
		inorderToSlice(r.right, list)
	}
}

func Intro() {

	//           a
	//         /   \
	//        b     c
	//       / \   /  \
	//      d   e f    g
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

	a := treeNode{val: "a"}
	b := treeNode{val: "b"}
	c := treeNode{val: "c"}
	d := treeNode{val: "d"}
	e := treeNode{val: "e"}
	f := treeNode{val: "f"}
	g := treeNode{val: "g"}

	fmt.Println(a, b, c, d, e, f)

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.left = &f
	c.right = &g

	var list []string

	prettyPrintBTree(&a)
	dfsRecursiveBinaryTree(&a)

	fmt.Println("here is inorder")
	inorderPrint(&a)
	list = []string{}
	inorderToSlice(&a, &list)
	fmt.Println(list)

	fmt.Println("here is preorder")
	preorderPrint(&a)
	list = []string{}
	preorderToSlice(&a, &list)
	fmt.Println(list)

	fmt.Println("here is postorder")
	postorderPrint(&a)
	list = []string{}
	postorderToSlice(&a, &list)
	fmt.Println(list)

}
