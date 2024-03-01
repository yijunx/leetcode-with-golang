package tutorials

import (
	"fmt"
)

func countConstruct(target string, materials []string) int {
	memo := map[string]int{}
	return countConstructRecursive(target, materials, memo)
}

func countConstructRecursive(target string, materials []string, memo map[string]int) int {
	if target == "" {
		// base case, we found 1 case
		return 1
	}
	if val, exists := memo[target]; exists {
		// well we did it before
		return val
	}
	for _, material := range materials {
		// fmt.Println("material:", material)
		if len(material) <= len(target) {
			// well obviously...
			if target[:len(material)] == material {
				// start with the world
				remainder := target[len(material):]
				// fmt.Println"remainder is", remainder)
				remainder_status := countConstructRecursive(remainder, materials, memo)
				memo[target] += remainder_status
			}
		}
	}
	return memo[target]
}

func CountConstruct() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(countConstruct("helloblueworld", []string{"hello", "blue", "world"}))
	fmt.Println(countConstruct("helloblueworld", []string{"hel", "llob", "world"}))
	fmt.Println(countConstruct("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	fmt.Println(countConstruct("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "enter", "ent", "ot", "o", "t"}))
}
