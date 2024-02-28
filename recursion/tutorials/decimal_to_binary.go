package tutorials

import (
	"fmt"
)

func dToB(i int) string {
	if i <= 1 {
		return fmt.Sprint(i)
	}

	return dToB(i/2) + fmt.Sprint(i%2)
}

func DecimalToBinary() {
	s := 127
	fmt.Println(dToB(s))
}
