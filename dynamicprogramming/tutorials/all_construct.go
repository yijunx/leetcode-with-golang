package tutorials

import (
	"fmt"
)

func allConstruct(target string, materials []string) [][]string {
	memo := map[string][][]string{}
	return allConstructRecursive(target, materials, memo)
}

func allConstructRecursive(target string, materials []string, memo map[string][][]string) [][]string {
	if target == "" {
		// base case..
		return [][]string{{}}
	}

	if val, exists := memo[target]; exists {
		return val
	}

	result := [][]string{}
	for _, material := range materials {
		if len(target) >= len(material) {
			// well other wise where is no way..
			if target[:len(material)] == material {
				// now it can be constructed
				remainder := target[len(material):]
				// fmt.Println(remainder)
				allWaysToMakeRemainder := allConstructRecursive(
					remainder, materials, memo,
				)

				for _, oneWayForRemainder := range allWaysToMakeRemainder {
					// do the crazy append here
					oneWayForTarget := []string{}
					oneWayForTarget = append(oneWayForTarget, oneWayForRemainder...)
					oneWayForTarget = append(oneWayForTarget, material)
					result = append(result, oneWayForTarget)
					// memo[target] = append(memo[target], oneWayForTarget)
				}
			}
		}
	}
	memo[target] = result
	return memo[target]
}

func AllConstruct() {
	fmt.Println(allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(allConstruct("helloblueworld", []string{"hello", "blue", "world"}))
	fmt.Println(allConstruct("helloblueworld", []string{"hel", "llob", "world"}))
	fmt.Println(allConstruct("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	fmt.Println(allConstruct("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	fmt.Println(allConstruct("enterapotentpot", []string{"a", "p", "enter", "ent", "ot", "o", "t"}))
}
