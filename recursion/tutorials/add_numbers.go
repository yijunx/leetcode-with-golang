package tutorials

import "fmt"

func addNumberTill(n int) int {
	if n == 0 {
		return 0
	}
	return n + addNumberTill(n-1)
}
func AddNumbers() {
	fmt.Println(addNumberTill(10))
}
