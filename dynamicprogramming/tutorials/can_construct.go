package tutorials

import "fmt"

func canConstruct(target string, materials []string) bool {
	memo := map[string]bool{}
	return canConstructRecursive(target, materials, memo)
}

func canConstructRecursive(target string, materials []string, memo map[string]bool) bool {
	if target == "" {
		return true
	}
	if can, exists := memo[target]; exists {
		return can
	}
	for _, material := range materials {
		// fmt.Println("material:", material)
		if len(material) <= len(target) {
			if target[:len(material)] == material {
				// fmt.Println("we are going in!!")
				remainder := target[len(material):]
				remainder_status := canConstructRecursive(remainder, materials, memo)
				memo[remainder] = remainder_status
				return remainder_status
			}
		}
	}
	memo[target] = false
	return false
}

func CanConstruct() {
	fmt.Println(canConstruct("helloblueworld", []string{"hello", "blue", "world"}))
	fmt.Println(canConstruct("helloblueworld", []string{"hel", "llob", "world"}))
	fmt.Println(canConstruct("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	fmt.Println(canConstruct("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))
}
