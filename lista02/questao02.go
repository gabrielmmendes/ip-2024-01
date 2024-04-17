package main

import "fmt"

func main() {
	var a, b, x float32 = 0, 0, 0

	fmt.Scan(&a, &b)

	for ; a < b; x++ {
		a *= 1.03
		fmt.Printf("%.0f \n", a)
	}

	fmt.Printf("ANOS = %.0f \n", x)
}
