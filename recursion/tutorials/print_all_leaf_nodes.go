package tutorials

import "fmt"

func printAllLeafNodes(node *treeNode) {
	// this is dfs
	if node != nil {
		if node.left == nil && node.right == nil {
			fmt.Println(node.val)
		}
		printAllLeafNodes(node.left)
		printAllLeafNodes(node.right)
	}
}

func PrintAllLeafNodes() {

	//          50
	//         /  \
	//        25   75
	//       /  \    \
	//      10   30   80
	//                  \
	//                  100

	a := treeNode{val: 50}
	b := treeNode{val: 25}
	c := treeNode{val: 75}
	d := treeNode{val: 10}
	e := treeNode{val: 30}
	f := treeNode{val: 80}
	g := treeNode{val: 100}

	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.right = &f
	f.right = &g

	printAllLeafNodes(&a)

}
