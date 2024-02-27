package tutorials

import "fmt"

// func sumOfLinkedListRecursive(n *node) any {
// 	if n == nil {
// 		return 0
// 	} else {
// 		return sumOfLinkedListRecursive(n.next) + n.val
// 	}
// }

func SumOfLinkedList() {
	numbers := []int{1, 2, 3, 4}
	a := sliceToLinkedList(numbers)

	//fmt.Println(sumOfLinkedListRecursive(a))
	fmt.Println(a)
}
