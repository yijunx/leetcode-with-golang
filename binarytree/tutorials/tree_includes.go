package tutorials

import "fmt"

func treeIncludesBFS(node *justANode, val string) bool {

	queue := []*justANode{}
	for {

		if node != nil {

			if node.val == val {
				return true
			}
			// add to queue
			queue = append(queue, node.left)
			queue = append(queue, node.right)

		}

		if len(queue) == 0 {
			break
		}

		node, queue = queue[0], queue[1:]
	}

	return false
}

func treeIncludesDFS(r *justANode, val string) bool {

	if r != nil {
		if r.val == val {
			// found it
			return true
		}
		return treeIncludesDFS(r.right, val) || treeIncludesDFS(r.left, val)
	}
	// if r is nil just false...
	return false
}

func TreeIncludes() {

	a := justANode{val: "a"}
	b := justANode{val: "b"}
	c := justANode{val: "c"}
	d := justANode{val: "d"}
	e := justANode{val: "e"}
	f := justANode{val: "f"}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	fmt.Println(treeIncludesBFS(&a, "e"))
	fmt.Println(treeIncludesDFS(&a, "g"))

}
