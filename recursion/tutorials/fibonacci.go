package tutorials

import "fmt"

func fiboSlow(n int) int {
	if n <= 1 {
		return n
	}
	return fiboSlow(n-1) + fiboSlow(n-2)
}

func fiboBetter(n int) int {
	if n <= 1 {
		return n
	}

	memo := [2]int{0, 1}

	return fiboWithMemory(n, memo)
}

func fiboWithMemory(n int, memo [2]int) int {
	// now n is at least 2

	result := memo[0] + memo[1]
	if n == 2 {
		return result
	}
	memo[0], memo[1] = memo[1], result
	return fiboWithMemory(n-1, memo)
}

func Fibonacci() {
	for i := 0; i < 10; i++ {
		fmt.Println(fiboSlow(i))
		fmt.Println(fiboBetter(i))
	}

}
