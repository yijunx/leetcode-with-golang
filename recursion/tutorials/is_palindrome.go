package tutorials

import (
	"fmt"
)

func isPal(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}
	// if len(s) == 2 {
	// 	return s[0] == s[1]
	// }

	if s[0] != s[len(s)-1] {
		return false
	}

	return isPal(s[1 : len(s)-1])
}

func IsPalindrome() {
	s := "dooggood"
	fmt.Println(s[1 : len(s)-1])
	fmt.Println(isPal(s))
}
