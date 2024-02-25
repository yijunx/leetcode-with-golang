package tutorials

import "fmt"

func treeSumBFS(node *numberNode) int {

	total := 0
	queue := []*numberNode{}

	for {

		if node != nil {
			total += node.val
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}

		if len(queue) == 0 {
			break
		}

		node, queue = queue[0], queue[1:]
	}
	return total
}

func treeSumDFS(node *numberNode) int {

	if node != nil {
		return node.val + treeSumDFS(node.left) + treeSumDFS(node.right)
	}
	return 0
}

func TreeSum() {

	a := numberNode{val: 1}
	b := numberNode{val: 2}
	c := numberNode{val: 3}
	d := numberNode{val: 4}
	e := numberNode{val: 5}
	f := numberNode{val: 6}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f

	fmt.Println(treeSumBFS(&a))
	fmt.Println(treeSumDFS(&a))

	const maxUInt = ^uint(0)
	fmt.Println(maxUInt)
	const maxInt = int(maxUInt >> 1)
	fmt.Println(maxInt)

}
