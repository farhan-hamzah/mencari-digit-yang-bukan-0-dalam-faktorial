
package main
import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(lastDigit(n))
}

func lastDigit(n int) int {
	if n == 0 {
		return 1
	}
	faktorial := 1
	for i := n; i >= 1; i-- {
		faktorial *= i
	}

	for faktorial > 0 {
		digit := faktorial % 10
		if digit != 0 {
			return digit
		}
		faktorial /= 10
	}

	return -1
}

