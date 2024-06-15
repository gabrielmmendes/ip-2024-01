package main

import "fmt"

func main() {
	var x, n int

	fmt.Scan(&x, &n)

	fmt.Println(elevado(x, n))
}

func elevado(n int, x int) int {
	if x <= 0 {
		return 1
	}
	return n * elevado(n, x-1)
}
