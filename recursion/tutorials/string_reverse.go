package tutorials

import (
	"fmt"
)

func stringReverse(s string) string {

	if s == "" {
		// ending condition
		return ""
	}

	return stringReverse(s[1:]) + s[0:1]
}

func runeReverse(r []rune) []rune {

	if len(r) == 0 {
		// ending condition
		return r
	}

	return append(runeReverse(r[1:]), r[0])
}

func StringReverse() {
	s := "你好good"
	fmt.Println(s)
	// fmt.Println(s[0:3]) //你
	// fmt.Println(s[1:3]) //dont know what
	// fmt.Println(s[3:6]) //好
	// fmt.Println(len(s)) //10

	// for i, v := range b {
	// 	fmt.Printf("index is %d, value is %v, type is %T\n", i, v, v)
	// }

	fmt.Println(stringReverse(s))
	// fmt.Println(runeReverse(b))

	r := []rune(s)
	// fmt.Println(r[0:1])
	fmt.Println(string(runeReverse(r)))
}
