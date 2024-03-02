package tutorials

import (
	"fmt"
)

func countConstruct(target string, materials []string) int {
	memo := map[string]int{}
	return countConstructRecursive(target, materials, memo)
}

func countConstructBFS(target string, materials []string) int {
	// memo := map[string]int{}
	queue := []string{}
	result := 0

	for {
		if target == "" {
			result += 1
		}
		// if _, exists := memo[target]; !exists {
		// 	memo[target] = true
		// }

		for _, material := range materials {
			if len(target) >= len(material) {
				if target[:len(material)] == material {
					remainder := target[len(material):]
					fmt.Println("remainder is", remainder)
					queue = append(queue, remainder)
				}
			}
		}

		if len(queue) == 0 {
			break
		}
		// fmt.Println(len(queue))
		target, queue = queue[0], queue[1:]
	}
	return result
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
	fmt.Println(countConstructBFS("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	// fmt.Println(countConstruct("helloblueworld", []string{"hello", "blue", "world"}))
	// fmt.Println(countConstruct("helloblueworld", []string{"hel", "llob", "world"}))
	// fmt.Println(countConstruct("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	// fmt.Println(countConstruct("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	// fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "enter", "ent", "ot", "o", "t"}))
}
