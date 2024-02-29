package tutorials

import "fmt"

func fiboWithMemo(n int, memo map[int]int) int {
	if n < 2 {
		memo[n] = n
		return n
	}
	if v, exists := memo[n]; exists {
		return v
	} else {
		v = fiboWithMemo(n-1, memo) + fiboWithMemo(n-2, memo)
		memo[n] = v
		return v
	}

	// return fiboWithMemo(n-1, memo) + fiboWithMemo(n-1, memo)
}

func easyFibo(n int) int {
	// with less memo consumption
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func Intro() {
	// just to show case Fibo with memo
	for i := 0; i < 10; i++ {
		fmt.Println(fiboWithMemo(i, map[int]int{}))
		fmt.Println(easyFibo(i))
	}
}
