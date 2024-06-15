package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(soma(x, len(x) - 1))
}

func soma(n []int, x int) int {
	if x < 0 {
		return 0
	}
	return n[x] + soma(n, x-1)
}
