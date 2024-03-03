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

func canContructTabu(target string, materials []string) bool {
	// dpResults := make([][]string, len(target)+1)
	dpConstructed := make([]string, len(target)+1)
	dpCan := make([]bool, len(target)+1)
	dpConstructed[0] = ""
	dpCan[0] = true

	for length, constructed := range dpConstructed {
		if !dpCan[length] {
			continue
		}

		for _, material := range materials {
			if len(material)+len(constructed) <= len(target) {
				// check if we have been here with the material
				if dpCan[len(material)+len(constructed)] {
					continue
				}
				// have not been here
				if constructed+material == target[:len(material)+len(constructed)] {
					dpConstructed[len(material)+len(constructed)] = constructed + material
					dpCan[len(material)+len(constructed)] = true
				}
			}
		}
	}
	return dpCan[len(target)]
}

func CanConstruct() {
	fmt.Println(canConstruct("helloblueworld", []string{"hello", "blue", "world"}))
	fmt.Println(canConstruct("helloblueworld", []string{"hel", "llob", "world"}))
	fmt.Println(canConstruct("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	fmt.Println(canConstruct("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))

	fmt.Println(canContructTabu("helloblueworld", []string{"hello", "blue", "world"}))
	fmt.Println(canContructTabu("helloblueworld", []string{"hel", "llob", "world"}))
	fmt.Println(canContructTabu("poiuthusadf", []string{"poi", "u", "adf", "th", "s"}))
	fmt.Println(canContructTabu("helloblueworldz", []string{"h", "e", "l", "o", "blue", "world"}))
	fmt.Println(canContructTabu("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}))
}
