package easy

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func symmetricTree(n *treeNode) bool {
	if n == nil {
		return true
	}
	return symmetricTreeRecursive(n.left, n.right)
}

func symmetricTreeRecursive(l *treeNode, r *treeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	fmt.Println(l.val, r.val)
	return l.val == r.val && symmetricTreeRecursive(l.left, r.right) && symmetricTreeRecursive(l.right, r.left)
}

func SymmetricTree() {

	a := treeNode{val: 1}
	b := treeNode{val: 2}
	c := treeNode{val: 2}
	d := treeNode{val: 3}
	e := treeNode{val: 4}
	f := treeNode{val: 4}
	g := treeNode{val: 3}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.left = &f
	c.right = &g

	fmt.Println(symmetricTree(&a))
}
